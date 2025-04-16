import { createContext, useContext, useState, useEffect } from 'react';
import { login as authLogin, register as authRegister } from '../services/authService';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const token = localStorage.getItem('token');
        console.log('Инициализация. Токен:', token); // <- Логируем
        setIsAuthenticated(!!token);
        setLoading(false);
    }, []);

    const login = async (credentials) => {
        try {
            setLoading(true);
            const token = await authLogin(credentials);
            console.log('Успешный вход. Токен:', token); // <- Логируем
            localStorage.setItem('token', token);
            setIsAuthenticated(true);
            return true;
        } catch (error) {
            console.error('Ошибка входа:', error); // <- Логируем ошибку
            localStorage.removeItem('token');
            setIsAuthenticated(false);
            throw error;
        } finally {
            setLoading(false);
        }
    };

    const register = async (userData) => {
        try {
            setLoading(true);
            await authRegister(userData);
            console.log('Регистрация успешна'); // <- Логируем
            return await login({
                email: userData.email,
                password: userData.password
            });
        } catch (error) {
            console.error('Ошибка регистрации:', error); // <- Логируем
            throw error;
        } finally {
            setLoading(false);
        }
    };

    const logout = () => {
        localStorage.removeItem('token');
        setIsAuthenticated(false);
    };

    console.log('AuthProvider render:', { isAuthenticated, loading }); // <- Логируем рендер

    return (
        <AuthContext.Provider value={{
            isAuthenticated,
            loading,
            login,
            register,
            logout
        }}>
            {!loading ? children : <div>Загрузка...</div>}
        </AuthContext.Provider>
    );
};

export const useAuth = () => useContext(AuthContext);