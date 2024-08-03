import React, { useState, useEffect } from 'react';
import TodoForm from './components/TodoForm';
import TodoList from './components/TodoList';
import { createTodo, getTodos, updateTodo, deleteTodo } from './api/todoApi';
import './App.css';

const App = () => {
    const [todos, setTodos] = useState([]);
    const categories = ['Work', 'Personal', 'Shopping'];

    const fetchTodos = async () => {
        const fetchedTodos = await getTodos();
        setTodos(fetchedTodos);
    };

    const handleAddTodo = async (newTodo) => {
        const createdTodo = await createTodo(newTodo);
        setTodos((prevTodos) => [...prevTodos, createdTodo]);
        fetchTodos();
    };

    const handleUpdate = async (id, updatedTodo) => {
        const updated = await updateTodo(id, updatedTodo);
        setTodos(todos.map((todo) => (todo.id === id ? updated : todo)));
    };

    const handleDelete = async (id) => {
        await deleteTodo(id);
        setTodos(todos.filter((todo) => todo.id !== id));
    };

    useEffect(() => {
        fetchTodos();
    }, []);

    return (
        <div className="app-container">
            <h1>Todo App</h1>
            <TodoForm onAdd={handleAddTodo} categories={categories} />
            <TodoList todos={todos} onUpdate={handleUpdate} onDelete={handleDelete} />
        </div>
    );
};

export default App;
