import { Todo } from "@/types/todo";

const API_URL = "http://localhost:8081/todos";

export async function fetchTodosServer(): Promise<Todo[]> {
  const res = await fetch(API_URL, {
  });
  if (!res.ok) {
    throw new Error(`Failed to fetch todos: ${res.status}`);
  }
  return res.json();
}

export async function createTodo(text: string): Promise<Todo> {
  const res = await fetch(API_URL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ text }),
  });
  if (!res.ok) {
    throw new Error("Failed to create todo");
  }
  return res.json();
}

export async function updateTodo(id: number, text: string, done: boolean): Promise<Todo> {
  const res = await fetch(`${API_URL}/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ text, done }),
  });
  if (!res.ok) {
    throw new Error("Failed to update todo");
  }
  return res.json();
}

export async function deleteTodo(id: number): Promise<void> {
  const res = await fetch(`${API_URL}/${id}`, { method: "DELETE" });
  if (!res.ok && res.status !== 204) {
    throw new Error("Failed to delete todo");
  }
}
