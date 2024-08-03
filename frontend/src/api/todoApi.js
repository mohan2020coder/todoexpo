import { apiGet, apiPost, apiPut, apiDelete } from './apiService';

const TODO_ENDPOINT = '/todos';

export const getTodos = () => apiGet(TODO_ENDPOINT);

export const createTodo = (todo) => apiPost(TODO_ENDPOINT, todo);

export const updateTodo = (id, todo) => apiPut(`${TODO_ENDPOINT}/${id}`, todo);

export const deleteTodo = (id) => apiDelete(`${TODO_ENDPOINT}/${id}`);
