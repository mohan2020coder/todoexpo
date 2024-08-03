import axios from 'axios';

// Create an instance of axios with default configurations
const apiClient = axios.create({
    baseURL: 'http://localhost:8080', // Base URL for API
    headers: {
        'Content-Type': 'application/json',
    },
});

const handleResponse = async (request) => {
    try {
        const response = await request;
        return response.data;
    } catch (error) {
        console.error("API Error:", error);
        throw error;
    }
};

export const apiGet = async (endpoint) => handleResponse(apiClient.get(endpoint));

export const apiPost = async (endpoint, data) => handleResponse(apiClient.post(endpoint, data));

export const apiPut = async (endpoint, data) => handleResponse(apiClient.put(endpoint, data));

export const apiDelete = async (endpoint) => handleResponse(apiClient.delete(endpoint));
