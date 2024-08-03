import React, { useState } from 'react';

const TodoForm = ({ onAdd, categories }) => {
    const [title, setTitle] = useState('');
    const [category, setCategory] = useState(categories[0]);
    const [dueDate, setDueDate] = useState('');
    const [reminder, setReminder] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        const newTodo = { title, category, dueDate, reminder };
        onAdd(newTodo);

        // Handle reminder if set
        if (reminder) {
            handleReminder(reminder);
        }

        // Reset form fields
        setTitle('');
        setCategory(categories[0]);
        setDueDate('');
        setReminder('');
    };

    const requestNotificationPermission = () => {
        if (Notification.permission === "default") {
            Notification.requestPermission().then(permission => {
                if (permission === "granted") {
                    console.log("Notification permission granted.");
                } else {
                    console.log("Notification permission denied.");
                }
            });
        }
    };
    const notifyUser = (message) => {
        if (Notification.permission === "granted") {
            new Notification(message);
        }
    };
    const handleReminder = (reminderDate) => {
        const now = new Date();
        const reminderTime = new Date(reminderDate).getTime();
    
        if (reminderTime > now.getTime()) {
            // Calculate the delay
            const delay = reminderTime - now.getTime();
    
            setTimeout(() => notifyUser('Reminder for your task!'), delay);
        }
    };
    

    // Request notification permission on component mount
    React.useEffect(() => {
        requestNotificationPermission();
    }, []);

    return (
        <form className="todo-form" onSubmit={handleSubmit}>
            <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Enter task title"
                required
            />
            <select
                value={category}
                onChange={(e) => setCategory(e.target.value)}
                required
            >
                {categories.map(cat => (
                    <option key={cat} value={cat}>{cat}</option>
                ))}
            </select>
            <input
                type="date"
                value={dueDate}
                onChange={(e) => setDueDate(e.target.value)}
                placeholder="Due date"
            />
            <input
                type="datetime-local"
                value={reminder}
                onChange={(e) => setReminder(e.target.value)}
                placeholder="Reminder"
            />
            <button type="submit">Add Todo</button>
        </form>
    );
};

export default TodoForm;
