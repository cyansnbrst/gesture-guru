import { useEffect, useState, useRef } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import { getGestureById } from '../../services/gestureService';
import { checkGesture } from '../../services/checkService';

import {
    Box,
    Typography,
    Container,
    CircularProgress,
    Button,
    IconButton,
    Stack,
    Paper,
    Chip,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    LinearProgress
} from '@mui/material';

import { ArrowBack, Share, CameraAlt, Check, Close } from '@mui/icons-material';

const GestureDetail = () => {
    const { id } = useParams();
    const { isAuthenticated } = useAuth();
    const navigate = useNavigate();

    const [gesture, setGesture] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const [cameraOpen, setCameraOpen] = useState(false);
    const [isChecking, setIsChecking] = useState(false);
    const [isRecording, setIsRecording] = useState(false);
    const [countdown, setCountdown] = useState(0);
    const [checkResult, setCheckResult] = useState(null);
    const [predictedGesture, setPredictedGesture] = useState(null);
    const [showCamera, setShowCamera] = useState(true); // Новое состояние для отображения камеры

    const videoRef = useRef(null);
    const streamRef = useRef(null);
    const mediaRecorderRef = useRef(null);
    const recordedChunksRef = useRef([]);
    const countdownIntervalRef = useRef(null);

    // Получение данных о жесте
    useEffect(() => {
        const fetchGesture = async () => {
            try {
                setLoading(true);
                const response = await getGestureById(id);
                setGesture(response.gesture);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        if (isAuthenticated) fetchGesture();
    }, [id, isAuthenticated]);

    // Очистка интервалов при размонтировании
    useEffect(() => {
        return () => {
            if (countdownIntervalRef.current) {
                clearInterval(countdownIntervalRef.current);
            }
            stopCamera();
        };
    }, []);

    // Включение камеры при открытии диалога
    useEffect(() => {
        if (cameraOpen && showCamera) {
            startCamera();
        } else {
            stopCamera();
        }
    }, [cameraOpen, showCamera]);

    // Включение камеры
    const startCamera = async () => {
        try {
            setCheckResult(null);
            setPredictedGesture(null);
            const stream = await navigator.mediaDevices.getUserMedia({
                video: {
                    width: 1280,
                    height: 720,
                    facingMode: 'user'
                }
            });
            streamRef.current = stream;
            if (videoRef.current) {
                videoRef.current.srcObject = stream;
            }
        } catch (err) {
            setError('Не удалось получить доступ к камере');
            console.error('Camera error:', err);
            setCameraOpen(false);
        }
    };

    // Выключение камеры
    const stopCamera = () => {
        if (streamRef.current) {
            streamRef.current.getTracks().forEach(track => track.stop());
            streamRef.current = null;
        }
        if (countdownIntervalRef.current) {
            clearInterval(countdownIntervalRef.current);
            countdownIntervalRef.current = null;
        }
        setIsRecording(false);
        setCountdown(0);
    };

    // Запуск записи
    const startRecording = async () => {
        if (!streamRef.current || !gesture || isRecording) return;

        recordedChunksRef.current = [];
        setIsRecording(true);
        setCountdown(2); // Устанавливаем таймер на 2 секунды

        // Запускаем таймер
        countdownIntervalRef.current = setInterval(() => {
            setCountdown(prev => {
                if (prev <= 1) {
                    clearInterval(countdownIntervalRef.current);
                    stopRecording();
                    return 0;
                }
                return prev - 1;
            });
        }, 1000);

        // Настройка MediaRecorder
        mediaRecorderRef.current = new MediaRecorder(streamRef.current, {
            mimeType: 'video/mp4',
            videoBitsPerSecond: 2500000
        });

        mediaRecorderRef.current.ondataavailable = (event) => {
            if (event.data.size > 0) {
                recordedChunksRef.current.push(event.data);
            }
        };

        mediaRecorderRef.current.start();
    };

    // Остановка записи и отправка на сервер
    const stopRecording = async () => {
        if (!mediaRecorderRef.current || mediaRecorderRef.current.state === 'inactive') return;

        mediaRecorderRef.current.onstop = async () => {
            try {
                setIsChecking(true);
                setShowCamera(false);
                const blob = new Blob(recordedChunksRef.current, { type: 'video/mp4' });

                const result = await checkGesture({
                    name: gesture.name,
                    video: blob
                });

                setCheckResult(result.is_correct);
                setPredictedGesture(result.predicted_gesture);
            } catch (err) {
                setError('Ошибка при проверке жеста');
                console.error('Check gesture error:', err);
            } finally {
                setIsChecking(false);
                setIsRecording(false);
            }
        };

        mediaRecorderRef.current.stop();
    };

    // Закрытие диалога
    const handleCloseDialog = () => {
        setCameraOpen(false);
        setCheckResult(null);
        setPredictedGesture(null);
        setShowCamera(true);
    };

    if (!isAuthenticated) {
        return (
            <Container maxWidth="sm" sx={{ py: 10, textAlign: 'center' }}>
                <Typography variant="h5" gutterBottom>
                    Для просмотра войдите в систему
                </Typography>
                <Button
                    variant="contained"
                    size="large"
                    onClick={() => navigate('/login')}
                    sx={{ mt: 3 }}
                >
                    Войти
                </Button>
            </Container>
        );
    }

    if (loading) {
        return (
            <Box sx={{ display: 'flex', justifyContent: 'center', py: 10 }}>
                <CircularProgress size={60} />
            </Box>
        );
    }

    if (error || !gesture) {
        return (
            <Container maxWidth="sm" sx={{ py: 10, textAlign: 'center' }}>
                <Typography variant="h5" color="error" gutterBottom>
                    {error || 'Жест не найден'}
                </Typography>
                <Button
                    variant="outlined"
                    size="large"
                    onClick={() => navigate('/gestures')}
                    sx={{ mt: 3 }}
                >
                    Назад к жестам
                </Button>
            </Container>
        );
    }

    const embedUrl = `https://www.youtube.com/embed/${gesture.videoUrl}`;

    return (
        <Container maxWidth="lg" sx={{ py: 4 }}>
            <Box sx={{ mb: 4 }}>
                <IconButton onClick={() => navigate(-1)} size="large">
                    <ArrowBack fontSize="large" />
                </IconButton>
            </Box>

            <Stack direction={{ xs: 'column', md: 'row' }} spacing={4}>
                {/* Видео-блок */}
                <Box sx={{ flex: 1 }}>
                    <Paper elevation={3} sx={{ borderRadius: 3, overflow: 'hidden' }}>
                        <Box
                            sx={{
                                position: 'relative',
                                paddingBottom: '56.25%',
                                height: 0,
                                overflow: 'hidden',
                                bgcolor: '#000'
                            }}
                        >
                            <iframe
                                src={embedUrl}
                                title={gesture.name}
                                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                                allowFullScreen
                                style={{
                                    position: 'absolute',
                                    top: 0,
                                    left: 0,
                                    width: '100%',
                                    height: '100%',
                                    border: 'none'
                                }}
                            />
                        </Box>
                    </Paper>

                    <Stack direction="row" spacing={2} sx={{ mt: 3 }}>
                        <Button
                            variant="contained"
                            startIcon={<CameraAlt />}
                            sx={{ flex: 1 }}
                            onClick={() => setCameraOpen(true)}
                        >
                            Проверить жест
                        </Button>
                        <Button variant="outlined" startIcon={<Share />} sx={{ flex: 1 }}>
                            Поделиться
                        </Button>
                    </Stack>
                </Box>

                {/* Информация о жесте */}
                <Box sx={{ width: { xs: '100%', md: '35%' } }}>
                    <Typography variant="h4" component="h1" gutterBottom>
                        {gesture.name}
                    </Typography>

                    <Chip
                        label={`Категория: ${gesture.categoryId}`}
                        variant="outlined"
                        sx={{ mb: 3 }}
                    />

                    <Typography variant="body1" paragraph sx={{ mb: 3 }}>
                        {gesture.description}
                    </Typography>

                    <Typography variant="caption" display="block" sx={{ mt: 2 }}>
                        Добавлено: {new Date(gesture.createdAt).toLocaleDateString()}
                    </Typography>
                </Box>
            </Stack>

            {/* Диалог камеры */}
            <Dialog
                open={cameraOpen}
                onClose={handleCloseDialog}
                maxWidth="md"
                fullWidth
                PaperProps={{
                    sx: {
                        borderRadius: 3
                    }
                }}
            >
                <DialogTitle sx={{ textAlign: 'center' }}>
                    Проверка жеста: {gesture.name}
                </DialogTitle>
                <DialogContent>
                    {showCamera ? (
                        <>
                            <Box sx={{
                                position: 'relative',
                                paddingBottom: '56.25%',
                                height: 0,
                                overflow: 'hidden',
                                bgcolor: '#000',
                                borderRadius: 2,
                                mb: 2
                            }}>
                                <video
                                    ref={videoRef}
                                    autoPlay
                                    playsInline
                                    muted
                                    style={{
                                        position: 'absolute',
                                        top: 0,
                                        left: 0,
                                        width: '100%',
                                        height: '100%',
                                        objectFit: 'cover',
                                        transform: 'scaleX(-1)'
                                    }}
                                />
                            </Box>

                            {isRecording && (
                                <Box sx={{ width: '100%', mb: 2 }}>
                                    <Typography variant="h6" align="center" gutterBottom>
                                        Запись через: {countdown} сек
                                    </Typography>
                                    <LinearProgress
                                        variant="determinate"
                                        value={(5 - countdown) * 20}
                                        sx={{ height: 10, borderRadius: 5 }}
                                    />
                                </Box>
                            )}
                        </>
                    ) : (
                        <Box sx={{
                            height: '400px',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center',
                            bgcolor: '#f5f5f5',
                            borderRadius: 2,
                            mb: 2
                        }}>
                            {isChecking ? (
                                <Box sx={{ textAlign: 'center' }}>
                                    <CircularProgress size={60} />
                                    <Typography variant="h6" sx={{ mt: 2 }}>
                                        Проверяем ваш жест...
                                    </Typography>
                                </Box>
                            ) : checkResult !== null && (
                                <Box sx={{ textAlign: 'center' }}>
                                    {checkResult ? (
                                        <>
                                            <Check color="success" sx={{ fontSize: 60 }} />
                                            <Typography variant="h5" color="success.main" sx={{ mt: 2 }}>
                                                Вы правильно показали жест!
                                            </Typography>
                                        </>
                                    ) : (
                                        <>
                                            <Close color="error" sx={{ fontSize: 60 }} />
                                            <Typography variant="h5" color="error.main" sx={{ mt: 2 }}>
                                                Вы показали: {predictedGesture || 'неизвестный жест'}
                                            </Typography>
                                            <Typography variant="body1" sx={{ mt: 2 }}>
                                                Попробуйте ещё раз
                                            </Typography>
                                        </>
                                    )}
                                </Box>
                            )}
                        </Box>
                    )}
                </DialogContent>
                <DialogActions sx={{ justifyContent: 'center', pb: 3 }}>
                    <Button
                        variant="outlined"
                        onClick={handleCloseDialog}
                        sx={{ mr: 2 }}
                    >
                        Закрыть
                    </Button>
                    {showCamera && (
                        <Button
                            variant="contained"
                            onClick={isRecording ? null : startRecording}
                            disabled={isChecking || isRecording}
                            startIcon={isChecking ? <CircularProgress size={20} color="inherit" /> : null}
                            sx={{ px: 4 }}
                            color={isRecording ? 'secondary' : 'primary'}
                        >
                            {isChecking ? 'Проверяем...' :
                                isRecording ? 'Идет запись...' : 'Начать запись (2 сек)'}
                        </Button>
                    )}
                    {!showCamera && !isChecking && (
                        <Button
                            variant="contained"
                            onClick={() => setShowCamera(true)}
                            sx={{ px: 4 }}
                        >
                            Попробовать снова
                        </Button>
                    )}
                </DialogActions>
            </Dialog>
        </Container>
    );
};

export default GestureDetail;
