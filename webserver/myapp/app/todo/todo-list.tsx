import styles from './page'
import UIkit from 'uikit'
import Icons from 'uikit/dist/js/uikit-icons'
import 'uikit/dist/css/uikit.css'
import 'uikit/dist/css/uikit.min.css'
UIkit.use(Icons)
import { Todo } from '../../typing';
import TodoLine from './todo-line'

const fetchTodos = async () => {
  const res = await fetch(`http://${process.env.APSERVER_HOST}:8080/api/todos`, {
    cache: "no-store",
  });
  const todos: Todo[] = await res.json();
  return todos
}

export default async function TodosList() {
  const todos = await fetchTodos();
  return (
    <>
      <table className="uk-table uk-table-middle uk-table-divider">
        <thead>
          <tr>
            <th className="uk-width-middle">Todo List</th>
            <th className="uk-width-small">Delete</th>
          </tr>
        </thead>
        <tbody>
          {
            todos.map((t) => (
              <tr key={t.id}>
                <TodoLine todo={t} />
              </tr>
            ))
          }
        </tbody>
      </table>
    </>
  );
}
