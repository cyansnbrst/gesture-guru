import { useEffect, useState } from 'react';
import { useAuth } from '../../context/AuthContext';
import { Link } from 'react-router-dom';
import {
    Box,
    Typography,
    Card,
    CardContent,
    Container,
    Button,
    Alert,
    CircularProgress
} from '@mui/material';
import { getGestures } from '../../services/gestureService';

const GesturesList = () => {
    const { isAuthenticated } = useAuth();
    const [gestures, setGestures] = useState([]);
    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    useEffect(() => {
        const fetchGestures = async () => {
            setIsLoading(true);
            try {
                const response = await getGestures();
                let gesturesData = [];
                if (Array.isArray(response)) {
                    gesturesData = response;
                } else if (response?.gestures && Array.isArray(response.gestures)) {
                    gesturesData = response.gestures;
                } else {
                    throw new Error('Некорректный формат данных');
                }
                setGestures(gesturesData);
                setError(null);
            } catch (err) {
                setError('Не удалось загрузить жесты. Попробуйте позже.');
                setGestures([]);
            } finally {
                setIsLoading(false);
            }
        };

        if (isAuthenticated) {
            fetchGestures();
        }
    }, [isAuthenticated]);

    if (!isAuthenticated) {
        return (
            <Container maxWidth="md" sx={{ mt: 8, textAlign: 'center' }}>
                <Typography variant="h5" sx={{ mb: 3 }}>
                    Войдите, чтобы просмотреть жесты
                </Typography>
                <Button
                    component={Link}
                    to="/login"
                    variant="contained"
                    sx={{
                        background: 'linear-gradient(to right, #4f46e5, #7c3aed)',
                        textTransform: 'none',
                        fontSize: '1rem',
                        px: 4,
                        py: 1.5
                    }}
                >
                    Войти
                </Button>
            </Container>
        );
    }

    return (
        <Box sx={{
            minHeight: '100vh',
            position: 'relative',
            overflow: 'hidden',
            backgroundColor: 'background.default',
            pt: 8,
            pb: 4
        }}>
            {/* Декоративные элементы */}
            <Box sx={{
                position: 'absolute',
                width: 200,
                height: 200,
                borderRadius: '50%',
                backgroundColor: 'primary.light',
                top: -50,
                left: -50,
                opacity: 0.6,
                filter: 'blur(40px)',
                zIndex: 0
            }} />
            <Box sx={{
                position: 'absolute',
                width: 200,
                height: 200,
                borderRadius: '50%',
                backgroundColor: 'secondary.light',
                bottom: -50,
                right: -50,
                opacity: 0.6,
                filter: 'blur(40px)',
                zIndex: 0
            }} />

            <Container maxWidth="lg" sx={{ position: 'relative', zIndex: 1 }}>
                <Typography variant="h4" sx={{ mb: 4, fontWeight: 700 }}>
                    Библиотека жестов
                </Typography>

                {error && (
                    <Alert severity="error" sx={{ mb: 3 }}>
                        {error}
                    </Alert>
                )}

                {isLoading ? (
                    <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
                        <CircularProgress />
                    </Box>
                ) : (
                    <>
                        {!gestures || gestures.length === 0 ? (
                            <Alert severity="info">Жесты не найдены</Alert>
                        ) : (
                            <Box
                                sx={{
                                    display: 'flex',
                                    flexWrap: 'wrap',
                                    justifyContent: 'center',
                                    gap: 4,
                                }}
                            >
                                {gestures.map((gesture) => (
                                    <Card
                                        key={gesture.id}
                                        sx={{
                                            width: '100%',
                                            maxWidth: 440,
                                            flex: '1 1 300px',
                                            display: 'flex',
                                            flexDirection: 'column',
                                            borderRadius: 2,
                                            overflow: 'hidden',
                                            boxShadow: 6,
                                            transition: 'transform 0.3s, box-shadow 0.3s',
                                            '&:hover': {
                                                transform: 'translateY(-5px)',
                                                boxShadow: 10,
                                            },
                                        }}
                                    >
                                        {/* Картинка - немного ниже по высоте */}
                                        <Box sx={{ position: 'relative', pt: '50%' }}>
                                            <img
                                                src={`https://img.youtube.com/vi/${gesture.videoUrl}/hqdefault.jpg`}
                                                alt={gesture.name}
                                                style={{
                                                    position: 'absolute',
                                                    top: 0,
                                                    left: 0,
                                                    width: '100%',
                                                    height: '100%',
                                                    objectFit: 'cover',
                                                }}
                                            />
                                        </Box>

                                        {/* Контент карточки */}
                                        <CardContent sx={{ display: 'flex', flexDirection: 'column', gap: 2, p: 3 }}>
                                            <Typography
                                                variant="caption"
                                                color="primary"
                                                sx={{
                                                    fontWeight: 700,
                                                    textTransform: 'uppercase',
                                                    letterSpacing: '0.06em',
                                                }}
                                            >
                                                Категория: {gesture.category_id}
                                            </Typography>

                                            {/* Название - крупнее */}
                                            <Typography
                                                variant="h5"
                                                fontWeight={700}
                                                sx={{ fontSize: '1.5rem', lineHeight: 1.3 }}
                                            >
                                                {gesture.name}
                                            </Typography>

                                            {/* Кнопка - капс и на всю ширину */}
                                            <Button
                                                component={Link}
                                                to={`/gestures/${gesture.id}`}
                                                variant="contained"
                                                fullWidth
                                                sx={{
                                                    mt: 1,
                                                    textTransform: 'uppercase',
                                                    px: 3,
                                                    py: 1.2,
                                                    borderRadius: 1,
                                                    fontWeight: 700,
                                                    letterSpacing: '0.05em',
                                                    background: '#6366f1',
                                                }}
                                            >
                                                Изучить
                                            </Button>
                                        </CardContent>
                                    </Card>

                                ))}
                            </Box>

                        )}
                    </>
                )}
            </Container>
        </Box >
    );
};

export default GesturesList;