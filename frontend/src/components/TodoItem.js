import React from 'react';
import { FaCheck, FaTimes, FaBell } from 'react-icons/fa';

const TodoItem = ({ todo, onUpdate, onDelete }) => {
    const handleToggleComplete = () => {
        onUpdate(todo.id, { ...todo, completed: !todo.completed });
    };

    const handleDelete = () => {
        onDelete(todo.id);
    };

    return (
        <div className={`todo-item ${todo.completed ? 'completed' : ''}`}>
            <div className="todo-item-content">
                <h3>{todo.title} - {todo.category}</h3>
                {todo.dueDate && <p>Due: {new Date(todo.dueDate).toLocaleDateString()}</p>}
                {todo.reminder && <p>Reminder: {new Date(todo.reminder).toLocaleString()}</p>}
            </div>
            <div className="todo-item-actions">
                <button className="toggle-complete-btn" onClick={handleToggleComplete}>
                    {todo.completed ? <FaTimes /> : <FaCheck />}
                </button>
                <button className="delete-btn" onClick={handleDelete}>
                    <FaTimes />
                </button>
                {todo.reminder && <FaBell />}
            </div>
        </div>
    );
};

export default TodoItem;
