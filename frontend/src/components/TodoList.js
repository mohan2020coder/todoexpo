import React,{useState} from 'react';
import TodoItem from './TodoItem';

const TodoList = ({todos, onUpdate, onDelete}) => {
   
    const [filter, setFilter] = useState('All');

   
   
    const filteredTodos = todos.filter(todo => filter === 'All' || todo.category === filter);

    return (
        <div className="todo-list">
            <div className="filter-buttons">
                <button onClick={() => setFilter('All')}>All</button>
                <button onClick={() => setFilter('Work')}>Work</button>
                <button onClick={() => setFilter('Personal')}>Personal</button>
                <button onClick={() => setFilter('Shopping')}>Shopping</button>
            </div>
            {filteredTodos.map((todo) => (
                <TodoItem
                    key={todo.id}
                    todo={todo}
                    onUpdate={onUpdate}
                    onDelete={onDelete}
                />
            ))}
        </div>
    );
};

export default TodoList;
