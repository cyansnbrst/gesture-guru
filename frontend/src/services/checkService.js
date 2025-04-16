export const checkGesture = async ({ name, video }) => {
    const formData = new FormData();
    formData.append("video", video, "gesture.mp4");
    formData.append("gesture", name);

    const response = await fetch("http://localhost:8000/check", {
        method: "POST",
        body: formData,
    });

    return await response.json();
};
