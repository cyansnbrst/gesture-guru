export const getItem = (key) => {
    try {
        return localStorage.getItem(key);
    } catch (error) {
        console.error('Error accessing localStorage', error);
        return null;
    }
};

export const setItem = (key, value) => {
    try {
        localStorage.setItem(key, value);
    } catch (error) {
        console.error('Error setting localStorage', error);
    }
};

export const removeItem = (key) => {
    try {
        localStorage.removeItem(key);
    } catch (error) {
        console.error('Error removing from localStorage', error);
    }
};