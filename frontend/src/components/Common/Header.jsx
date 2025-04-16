import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import {
    AppBar,
    Toolbar,
    Button,
    Box,
    Avatar,
    Menu,
    MenuItem,
    IconButton,
    Typography
} from '@mui/material';
import {
    AccountCircle,
    ExitToApp,
    Menu as MenuIcon
} from '@mui/icons-material';

const Header = () => {
    const { logout } = useAuth();
    const navigate = useNavigate();
    const [anchorEl, setAnchorEl] = useState(null);
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    // Проверяем наличие токена в localStorage при монтировании компонента
    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            setIsAuthenticated(true); // Пользователь авторизован
        } else {
            setIsAuthenticated(false); // Пользователь не авторизован
        }
    }, []);

    const handleMenu = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    const handleLogout = () => {
        logout();
        localStorage.removeItem('token'); // Удаляем токен из localStorage
        setIsAuthenticated(false); // Пользователь выходит
        handleClose();
        navigate('/login');
    };

    return (
        <AppBar
            position="static"
            elevation={0}
            sx={{
                backgroundColor: 'background.paper',
                color: 'text.primary',
                borderBottom: '1px solid',
                borderColor: 'divider',
                py: 1,
                position: 'relative',
                '&:before': {
                    content: '""',
                    position: 'absolute',
                    width: 200,
                    height: 200,
                    borderRadius: '50%',
                    backgroundColor: 'primary.light',
                    top: -100,
                    left: -100,
                    opacity: 0.3,
                    filter: 'blur(40px)',
                    zIndex: 0
                },
                '&:after': {
                    content: '""',
                    position: 'absolute',
                    width: 200,
                    height: 200,
                    borderRadius: '50%',
                    backgroundColor: 'secondary.light',
                    bottom: -100,
                    right: -100,
                    opacity: 0.3,
                    filter: 'blur(40px)',
                    zIndex: 0
                }
            }}
        >
            <Toolbar sx={{
                justifyContent: 'space-between',
                position: 'relative',
                zIndex: 1
            }}>
                {/* Логотип и название */}
                <Box
                    component={Link}
                    to="/"
                    sx={{
                        display: 'flex',
                        alignItems: 'center',
                        textDecoration: 'none',
                        color: 'inherit',
                        '&:hover': {
                            opacity: 0.8
                        }
                    }}
                >
                    <Avatar
                        sx={{
                            mr: 2,
                            bgcolor: 'primary.main',
                            color: 'primary.contrastText',
                            fontWeight: 'bold'
                        }}
                    >
                        GG
                    </Avatar>
                    <Typography
                        variant="h6"
                        component="span"
                        sx={{
                            fontWeight: 700,
                            background: 'linear-gradient(to right, #4f46e5, #7c3aed)',
                            WebkitBackgroundClip: 'text',
                            backgroundClip: 'text',
                            color: 'transparent'
                        }}
                    >
                        GestureGuru
                    </Typography>
                </Box>

                {/* Навигация для десктопа */}
                <Box sx={{
                    display: { xs: 'none', md: 'flex' },
                    gap: 2,
                    alignItems: 'center'
                }}>
                    {isAuthenticated ? (
                        <>
                            <Button
                                component={Link}
                                to="/gestures"
                                color="inherit"
                                sx={{
                                    textTransform: 'none',
                                    fontSize: '1rem',
                                    '&:hover': {
                                        backgroundColor: 'action.hover'
                                    }
                                }}
                            >
                                Библиотека
                            </Button>
                            <IconButton
                                size="large"
                                onClick={handleMenu}
                                color="inherit"
                                sx={{
                                    '&:hover': {
                                        backgroundColor: 'action.hover'
                                    }
                                }}
                            >
                                <AccountCircle fontSize="large" />
                            </IconButton>
                            <Menu
                                anchorEl={anchorEl}
                                open={Boolean(anchorEl)}
                                onClose={handleClose}
                                PaperProps={{
                                    sx: {
                                        mt: 1.5,
                                        minWidth: 180,
                                        boxShadow: 2,
                                        borderRadius: 1.5
                                    }
                                }}
                            >
                                <MenuItem
                                    onClick={handleLogout}
                                    sx={{
                                        py: 1.5,
                                        '&:hover': {
                                            backgroundColor: 'action.hover'
                                        }
                                    }}
                                >
                                    <ExitToApp sx={{ mr: 1.5, color: 'error.main' }} />
                                    <Typography>Выйти</Typography>
                                </MenuItem>
                            </Menu>
                        </>
                    ) : (
                        <>
                            <Button
                                component={Link}
                                to="/login"
                                color="inherit"
                                sx={{
                                    textTransform: 'none',
                                    fontSize: '1rem',
                                    '&:hover': {
                                        backgroundColor: 'action.hover'
                                    }
                                }}
                            >
                                Войти
                            </Button>
                            <Button
                                component={Link}
                                to="/register"
                                variant="contained"
                                sx={{
                                    textTransform: 'none',
                                    fontSize: '1rem',
                                    background: 'linear-gradient(to right, #4f46e5, #7c3aed)',
                                    '&:hover': {
                                        opacity: 0.9,
                                        background: 'linear-gradient(to right, #4f46e5, #7c3aed)'
                                    }
                                }}
                            >
                                Регистрация
                            </Button>
                        </>
                    )}
                </Box>

                {/* Мобильное меню */}
                <Box sx={{ display: { xs: 'flex', md: 'none' } }}>
                    <IconButton
                        size="large"
                        color="inherit"
                        onClick={handleMenu}
                        sx={{
                            '&:hover': {
                                backgroundColor: 'action.hover'
                            }
                        }}
                    >
                        <MenuIcon />
                    </IconButton>
                    <Menu
                        anchorEl={anchorEl}
                        open={Boolean(anchorEl)}
                        onClose={handleClose}
                        PaperProps={{
                            sx: {
                                mt: 1.5,
                                minWidth: 180,
                                boxShadow: 2,
                                borderRadius: 1.5
                            }
                        }}
                    >
                        {isAuthenticated ? (
                            [
                                <MenuItem
                                    key="logout"
                                    onClick={handleLogout}
                                    sx={{
                                        py: 1.5,
                                        '&:hover': {
                                            backgroundColor: 'action.hover'
                                        }
                                    }}
                                >
                                    <ExitToApp sx={{ mr: 1.5, color: 'error.main' }} />
                                    <Typography>Выйти</Typography>
                                </MenuItem>
                            ]
                        ) : (
                            [
                                <MenuItem
                                    key="login"
                                    component={Link}
                                    to="/login"
                                    onClick={handleClose}
                                    sx={{
                                        py: 1.5,
                                        '&:hover': {
                                            backgroundColor: 'action.hover'
                                        }
                                    }}
                                >
                                    Войти
                                </MenuItem>,
                                <MenuItem
                                    key="register"
                                    component={Link}
                                    to="/register"
                                    onClick={handleClose}
                                    sx={{
                                        py: 1.5,
                                        '&:hover': {
                                            backgroundColor: 'action.hover'
                                        },
                                        color: 'primary.main',
                                        fontWeight: 'medium'
                                    }}
                                >
                                    Регистрация
                                </MenuItem>
                            ]
                        )}
                    </Menu>
                </Box>
            </Toolbar>
        </AppBar>
    );
};

export default Header;
