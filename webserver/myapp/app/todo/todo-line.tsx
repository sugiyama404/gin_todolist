"use client";

import { useRouter } from "next/navigation";

async function deleteTodo(id, refresh) {
  await fetch(`/api/todos/${id}`, {
    method: "DELETE",
  });

  refresh();
}

export default function TodoLine({ todo }) {
  const router = useRouter();

  return (
    <>
      <td className="uk-width-middle">
        {todo.content}
      </td>
      <td className="uk-width-small">
        <button className="uk-button uk-button-default" type="button"
          onClick={() => deleteTodo(todo.id, router.refresh)}>
          Delete
        </button>
      </td>
    </>
  );
}
