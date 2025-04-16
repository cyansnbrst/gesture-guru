import api from '../utils/api';

export const getGestures = async () => {
    const response = await api.get('/gestures');
    return response.data;
};

export const getGestureById = async (id) => {
    const response = await api.get(`/gestures/${id}`);
    return response.data;
};