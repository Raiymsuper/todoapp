import { useState, useEffect } from "react";
import { AddTask, GetTasks, ToggleTaskStatus, DeleteTask } from "../wailsjs/go/main/App";

function App() {
    const [tasks, setTasks] = useState([]);
    const [newTask, setNewTask] = useState("");

    // Загружаем задачи при загрузке страницы
    useEffect(() => {
        loadTasks();
    }, []);

    function addTask() {
        const input = document.getElementById("taskInput");
        const text = input.value.trim();
    
        if (!text) return;
    
        AddTask(text);
        input.value = "";
    
        // Refresh task list after adding
        loadTasks();
    }
    
    async function loadTasks() {
        const tasks = GetTasks();
        const taskList = document.getElementById("taskList");
        taskList.innerHTML = "";  // Clear old list
    
        tasks.forEach(task => {
            const li = document.createElement("li");
            li.textContent = task.text;
            taskList.appendChild(li);
        });
    }
    
    // Load tasks when the page opens
    window.onload = loadTasks;
    

    const handleToggleTask = async (id) => {
        await ToggleTaskStatus(id);
        loadTasks();
    };

    const handleDeleteTask = async (id) => {
        await DeleteTask(id);
        loadTasks();
    };

    return (
        <div>
            <h1>To-Do List</h1>
            <input
                type="text"
                value={newTask}
                onChange={(e) => setNewTask(e.target.value)}
                placeholder="Введите задачу..."
            />
            <button onClick={addTask}>Добавить</button>

            <ul>
                {tasks.map((task) => (
                    <li key={task.id}>
                        <span
                            style={{ textDecoration: task.done ? "line-through" : "none", cursor: "pointer" }}
                            onClick={() => handleToggleTask(task.id)}
                        >
                            {task.text}
                        </span>
                        <button onClick={() => handleDeleteTask(task.id)}>Удалить</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;