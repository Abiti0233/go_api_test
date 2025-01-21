import { fetchTodosServer } from "@/api/todos";
import { Todo } from "@/types/todo";
import TodoList from "@/components/todo-list";

export default async function HomePage() {
  let todos: Todo[] = [];

  todos = await fetchTodosServer();

  return (
    <main className="max-w-md mx-auto pt-12">
      <h1 className="text-center mb-4 font-bold text-2xl">TODOアプリ</h1>
      <TodoList initialTodos={todos} />
    </main>
  );
}
