import api from '../utils/api';

export const login = async (credentials) => {
    const response = await api.post('/login', credentials);
    return response.data.token;
};

export const register = async (userData) => {
    await api.post('/register', userData);
};