import React, { useState } from 'react';

const TodoForm = ({ onAdd, categories }) => {
    const [title, setTitle] = useState('');
    const [category, setCategory] = useState(categories[0]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const newTodo = { title, category, completed: false };
        await onAdd(newTodo);
        setTitle('');
        setCategory(categories[0]);
    };

    return (
        <form onSubmit={handleSubmit} className="todo-form">
            <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Enter todo title"
                required
            />
            <select
                className='category-select'
                value={category}
                onChange={(e) => setCategory(e.target.value)}
                required
            >
                {categories.map((cat, index) => (
                    <option key={index} value={cat}>
                        {cat}
                    </option>
                ))}
            </select>
            <br />
            <button type="submit">Add Todo</button>
        </form>
    );
};

export default TodoForm;
