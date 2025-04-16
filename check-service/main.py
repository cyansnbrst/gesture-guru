import argparse
import sys
import time
from collections import deque
from multiprocessing import Manager, Process, Value
from typing import Optional, Tuple, List
import tempfile

import onnxruntime as ort
from loguru import logger
import cv2
import numpy as np
from omegaconf import OmegaConf
from fastapi import FastAPI, UploadFile, File, Form, HTTPException
from fastapi.middleware.cors import CORSMiddleware

ort.set_default_logger_severity(4)  # NOQA
logger.add(sys.stdout, format="{level} | {message}")  # NOQA
logger.remove(0)  # NOQA

from constants import classes

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"], 
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class BaseRecognition:
    def __init__(self, model_path: str, tensors_list, prediction_list, verbose):
        self.verbose = verbose
        self.started = None
        self.output_names = None
        self.input_shape = None
        self.input_name = None
        self.session = None
        self.model_path = model_path
        self.window_size = None
        self.tensors_list = tensors_list
        self.prediction_list = prediction_list

    def clear_tensors(self):
        """Clear the list of tensors."""
        for _ in range(self.window_size):
            self.tensors_list.pop(0)

    def run(self):
        """Run the recognition model."""
        if self.session is None:
            self.session = ort.InferenceSession(self.model_path)
            self.input_name = self.session.get_inputs()[0].name
            self.input_shape = self.session.get_inputs()[0].shape
            self.window_size = self.input_shape[3]
            self.output_names = [output.name for output in self.session.get_outputs()]

        if len(self.tensors_list) >= self.input_shape[3]:
            input_tensor = np.stack(self.tensors_list[: self.window_size], axis=1)[None][None]
            st = time.time()
            outputs = self.session.run(self.output_names, {self.input_name: input_tensor.astype(np.float32)})[0]
            et = round(time.time() - st, 3)
            gloss = str(classes[outputs.argmax()])
            if gloss != self.prediction_list[-1] and len(self.prediction_list):
                if gloss != "---":
                    self.prediction_list.append(gloss)
            self.clear_tensors()
            if self.verbose:
                logger.info(f"- Prediction time {et}, new gloss: {gloss}")
                logger.info(f" --- {len(self.tensors_list)} frames in queue")

    def kill(self):
        pass


class Recognition(BaseRecognition):
    def __init__(self, model_path: str, tensors_list: list, prediction_list: list, verbose: bool):
        super().__init__(
            model_path=model_path, tensors_list=tensors_list, prediction_list=prediction_list, verbose=verbose
        )
        self.started = True

    def start(self):
        self.run()


class RecognitionMP(Process, BaseRecognition):
    def __init__(self, model_path: str, tensors_list, prediction_list, verbose):
        super().__init__()
        BaseRecognition.__init__(
            self, model_path=model_path, tensors_list=tensors_list, prediction_list=prediction_list, verbose=verbose
        )
        self.started = Value("i", False)

    def run(self):
        while True:
            BaseRecognition.run(self)
            self.started = True


class VideoProcessor:
    STACK_SIZE = 6

    def __init__(
            self,
            model_path: str,
            config: OmegaConf = None,
            mp: bool = False,
            verbose: bool = False,
            length: int = STACK_SIZE,
    ) -> None:
        self.multiprocess = mp
        self.manager = Manager() if self.multiprocess else None
        self.tensors_list = self.manager.list() if self.multiprocess else []
        self.prediction_list = self.manager.list() if self.multiprocess else []
        self.prediction_list.append("---")
        self.frame_counter = 0
        self.frame_interval = config.frame_interval
        self.length = length
        self.mean = config.mean
        self.std = config.std
        
        if self.multiprocess:
            self.recognizer = RecognitionMP(model_path, self.tensors_list, self.prediction_list, verbose)
        else:
            self.recognizer = Recognition(model_path, self.tensors_list, self.prediction_list, verbose)

    @staticmethod
    def resize(im, new_shape=(224, 224)):
        shape = im.shape[:2]
        if isinstance(new_shape, int):
            new_shape = (new_shape, new_shape)

        r = min(new_shape[0] / shape[0], new_shape[1] / shape[1])
        new_unpad = int(round(shape[1] * r)), int(round(shape[0] * r))
        dw, dh = new_shape[1] - new_unpad[0], new_shape[0] - new_unpad[1]
        dw /= 2
        dh /= 2

        if shape[::-1] != new_unpad:
            im = cv2.resize(im, new_unpad, interpolation=cv2.INTER_LINEAR)
        
        top, bottom = int(round(dh - 0.1)), int(round(dh + 0.1))
        left, right = int(round(dw - 0.1)), int(round(dw + 0.1))
        return cv2.copyMakeBorder(im, top, bottom, left, right, cv2.BORDER_CONSTANT, value=(114, 114, 114))

    def process_video(self, video_bytes: bytes) -> List[str]:
        """Process video and return predictions"""
        with tempfile.NamedTemporaryFile(suffix=".mp4") as tmp_file:
            tmp_file.write(video_bytes)
            cap = cv2.VideoCapture(tmp_file.name)
        
        predictions = []
        
        while cap.isOpened():
            ret, frame = cap.read()
            if not ret: 
                break
            
            self.frame_counter += 1
            if self.frame_counter % self.frame_interval == 0:
                frame = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
                frame = self.resize(frame, (224, 224))
                frame = (frame - self.mean) / self.std
                frame = np.transpose(frame, [2, 0, 1])
                self.tensors_list.append(frame)
                
                if not self.multiprocess:
                    self.recognizer.start()
                
                if len(self.prediction_list) > 1:  # Skip initial "---"
                    predictions.append(self.prediction_list[-1])
        
        cap.release()
        return predictions


def parse_arguments(params: Optional[Tuple] = None) -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Sign language recognition API")
    parser.add_argument("-p", "--config", required=True, type=str, help="Path to config")
    parser.add_argument("--mp", required=False, action="store_true", help="Enable multiprocessing")
    parser.add_argument("-v", "--verbose", required=False, action="store_true", help="Enable logging")
    parser.add_argument("-l", "--length", required=False, type=int, default=4, help="Deque length for predictions")
    
    known_args, _ = parser.parse_known_args(params)
    return known_args


# Initialize video processor with config
args = parse_arguments()
conf = OmegaConf.load(args.config)
video_processor = VideoProcessor(conf.model_path, conf, args.mp, args.verbose, args.length)


@app.post("/check")
async def check(
    video: UploadFile = File(...),
    gesture: str = Form(...)
):
    try:
        video_bytes = await video.read()
        predictions = video_processor.process_video(video_bytes)
        
        if not predictions:
            raise HTTPException(
                status_code=400,
                detail=f"Video too short or no predictions made. Need at least {video_processor.window_size} frames."
            )
        
        # Get the most frequent prediction
        predicted_gesture = max(set(predictions), key=predictions.count)
        is_correct = predicted_gesture.lower() == gesture.lower()
        
        return {
            "expected_gesture": gesture,
            "predicted_gesture": predicted_gesture,
            "is_correct": is_correct,
        }
        
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)