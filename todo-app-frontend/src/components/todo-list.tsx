"use client";

import { useState } from "react";
import { Todo } from "@/types/todo";
import { createTodo, updateTodo, deleteTodo } from "@/api/todos";

interface initialTodoProps {
  initialTodos: Todo[];
}

export default function TodoList({ initialTodos }: initialTodoProps) {
  const [todos, setTodos] = useState<Todo[]>(initialTodos);
  const [text, setText] = useState("");

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!text.trim()) return;

    try {
      const newTodo = await createTodo(text);
      setTodos([...todos, newTodo]);
      setText("");
    } catch (err) {
      console.error("Create failed:", err);
    }
  };

  const handleToggleDone = async (todo: Todo) => {
    try {
      const updated = await updateTodo(todo.ID, todo.Text, !todo.Done);
      setTodos(todos.map((t) => (t.ID === todo.ID ? updated : t)));
    } catch (err) {
      console.error("Update failed:", err);
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteTodo(id);
      setTodos(todos.filter((t) => t.ID !== id));
    } catch (err) {
      console.error("Delete failed:", err);
    }
  };

  return (
    <div>
      <form onSubmit={handleCreate} className="mb-4 flex justify-center">
        <input
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="新しいTODOを入力"
          className="p-2 border border-solid border-gray-500 rounded-md mr-2"
        />
        <button type="submit" className="bg-gray-100 py-2 px-6 border border-solid border-gray-500 rounded-md hover:bg-gray-300">追加</button>
      </form>
      <ul className="">
        {todos.map((todo) => (
          <li key={todo.ID} style={{ marginBottom: "0.5rem" }}>
            <div className="flex justify-between items-center max-w-[250px] mx-auto">
              <label
                className={`cursor-pointer ${todo.Done ? "line-through" : ""}`}
              >
                <input
                  type="checkbox"
                  checked={todo.Done}
                  onChange={() => handleToggleDone(todo)}
                  className="mr-2"
                />
                {todo.Text}
              </label>
              <button
                onClick={() => handleDelete(todo.ID)}
                className="text-red-500 hover:text-red-900 text-sm"
              >
                削除
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}
