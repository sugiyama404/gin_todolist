"use client";
import { useRouter } from "next/navigation";
import { useState } from "react";

async function addTodo(content, refresh) {

  const formData = new FormData();
  formData.append('content', content);

  await fetch("/api/todos", {
    method: "POST",
    body: formData,
  });

  refresh();
}

export default function AddNewTodo() {
  const router = useRouter();
  let [content, setContent] = useState("");
  return (
    <div>
      <form className="uk-form-horizontal uk-margin-large">
        <div className="uk-margin uk-flex">
          <input
            className="uk-input"
            type="text"
            onChange={(e) => setContent(e.target.value)}
            value={content}
          />
          <button
            className="uk-button uk-button-default"
            onClick={async () => {
              await addTodo(content, router.refresh);
              setContent("");
            }}
          >
            Add
          </button>
        </div>
      </form>
    </div>
  );
}
