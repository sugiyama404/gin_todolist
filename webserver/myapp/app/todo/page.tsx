import styles from './page.module.css'
import UIkit from 'uikit'
import Icons from 'uikit/dist/js/uikit-icons'
import 'uikit/dist/css/uikit.css'
import 'uikit/dist/css/uikit.min.css'
UIkit.use(Icons)
import TodosList from './todo-list'
import TodosForm from './todo-form'



export default async function Todo() {
  return (
    <>
      <div className="uk-container uk-container-xsmall">
        <h1>Todo</h1>
        {/* @ts-ignore */}
        <TodosForm />
        <TodosList />
      </div>
    </>
  );
}
