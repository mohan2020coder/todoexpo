import React from 'react';
import { FaCheck, FaTimes } from 'react-icons/fa'; // Import icons from react-icons
import './TodoItem.css'; // Import the CSS file

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
                <h3>{todo.title}</h3>
            </div>
            <div className="todo-item-actions">
                <button className="toggle-complete-btn" onClick={handleToggleComplete}>
                    {todo.completed ? <FaTimes /> : <FaCheck />}
                </button>
                <button className="delete-btn" onClick={handleDelete}>
                    <FaTimes />
                </button>
            </div>
        </div>
    );
};

export default TodoItem;
