import axios from 'axios';

//const baseURL = 'https://ethoe.dev/api'; // Replace with the URL of your Golang backend API
//const baseURL = 'http://localhost:5050/api'
const baseURL = process.env.REACT_APP_API_URL || 'https://ethoe.dev/api';


const apiClient = axios.create({
    baseURL,
    headers: {
        'Content-Type': 'application/json',
    },
});

export const get = async (url) => {
    try {
        const response = await apiClient.get(url);
        return response.data;
    } catch (error) {
        console.error('Error fetching data:', error);
        throw error;
    }
};

export const post = async (url, data) => {
    try {
        const response = await apiClient.post(url, data);
        return response.data;
    } catch (error) {
        console.error('Error creating data:', error);
        throw error;
    }
};

export const postFormData = async (url, formData) => {
    try {
        const response = await apiClient.post(url, formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        });
        return response.data;
    } catch (error) {
        console.error('Error creating data:', error);
        throw error;
    }
};

export const put = async (url, data) => {
    try {
        const response = await apiClient.put(url, data);
        return response.data;
    } catch (error) {
        console.error('Error updating data:', error);
        throw error;
    }
};

export const remove = async (url) => {
    try {
        const response = await apiClient.delete(url);
        return response.data;
    } catch (error) {
        console.error('Error deleting data:', error);
        throw error;
    }
};

export default apiClient;
