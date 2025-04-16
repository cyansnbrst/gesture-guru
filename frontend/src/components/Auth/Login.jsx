import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import {
    Box,
    TextField,
    Link,
    Alert,
    Fade,
    InputAdornment,
    IconButton,
    Typography,
    Button,
    Paper
} from '@mui/material';
import {
    Email as EmailIcon,
    Lock as LockIcon,
    Visibility,
    VisibilityOff
} from '@mui/icons-material';

const Login = () => {
    const { login } = useAuth();
    const navigate = useNavigate();
    const [error, setError] = useState('');
    const [showPassword, setShowPassword] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const credentials = {
            email: formData.get('email'),
            password: formData.get('password')
        };

        try {
            await login(credentials);
            navigate('/gestures');
        } catch (error) {
            setError('Неверные учетные данные');
        }
    };

    return (
        <Box sx={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            minHeight: '100vh',
            position: 'relative',
            overflow: 'hidden',
            backgroundColor: 'background.default'
        }}>
            <Box sx={{
                position: 'absolute',
                width: 200,
                height: 200,
                borderRadius: '50%',
                backgroundColor: 'primary.light',
                top: -50,
                left: -50,
                opacity: 0.6,
                filter: 'blur(40px)'
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
                filter: 'blur(40px)'
            }} />

            <Paper component="form" onSubmit={handleSubmit} sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                padding: 4,
                width: '100%',
                maxWidth: 400,
                position: 'relative',
                zIndex: 1,
                backgroundColor: 'background.paper',
                borderRadius: 2,
                boxShadow: 3
            }}>
                <Typography variant="h4" sx={{ mb: 3 }}>
                    Добро пожаловать
                </Typography>

                {error && (
                    <Fade in={!!error}>
                        <Alert severity="error" sx={{ width: '100%', mb: 3 }}>
                            {error}
                        </Alert>
                    </Fade>
                )}

                <TextField
                    fullWidth
                    margin="normal"
                    required
                    id="email"
                    label="Электронная почта"
                    name="email"
                    autoComplete="email"
                    InputProps={{
                        startAdornment: (
                            <InputAdornment position="start">
                                <EmailIcon color="action" />
                            </InputAdornment>
                        ),
                    }}
                />

                <TextField
                    fullWidth
                    margin="normal"
                    required
                    name="password"
                    label="Пароль"
                    type={showPassword ? 'text' : 'password'}
                    id="password"
                    InputProps={{
                        startAdornment: (
                            <InputAdornment position="start">
                                <LockIcon color="action" />
                            </InputAdornment>
                        ),
                        endAdornment: (
                            <InputAdornment position="end">
                                <IconButton
                                    onClick={() => setShowPassword(!showPassword)}
                                    edge="end"
                                >
                                    {showPassword ? <VisibilityOff /> : <Visibility />}
                                </IconButton>
                            </InputAdornment>
                        ),
                    }}
                />

                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    sx={{
                        mt: 3,
                        mb: 2,
                        py: 1.5
                    }}
                >
                    Войти
                </Button>

                <Box sx={{ mt: 2 }}>
                    <Link href="/register" color="primary" underline="hover">
                        Нет аккаунта? Зарегистрироваться
                    </Link>
                </Box>
            </Paper>
        </Box>
    );
};

export default Login;
