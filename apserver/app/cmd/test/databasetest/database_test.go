package databasetest

import (
	"app/cmd/infrastructure/dao"
	"testing"

	"app/cmd/interface/repository"

	"github.com/stretchr/testify/assert"
)

func TestTodoSave(t *testing.T) {
	conn := dao.ConnectDB()
	todo_repository := &repository.TodoRepository{Conn: conn}
	content := "test"
	todo, _ := todo_repository.Save(content)

	assert.Equal(t, todo.Content, "test")
}

func TestTodoFindAll(t *testing.T) {
	conn := dao.ConnectDB()
	todo_repository := &repository.TodoRepository{Conn: conn}
	todos, _ := todo_repository.FindAll()

	assert.NotEqual(t, len(todos), 0)
}

func TestTodoFindById(t *testing.T) {
	conn := dao.ConnectDB()
	todo_repository := &repository.TodoRepository{Conn: conn}
	todos, _ := todo_repository.FindAll()
	id := todos[len(todos)-1].ID
	todo, _ := todo_repository.FindById(id)

	assert.Equal(t, todo.ID, id)
}

func TestTodoUpdateById(t *testing.T) {
	conn := dao.ConnectDB()
	todo_repository := &repository.TodoRepository{Conn: conn}
	todos, _ := todo_repository.FindAll()
	id := todos[len(todos)-1].ID
	content := "newtest"
	todo, _ := todo_repository.UpdateById(id, content)

	assert.Equal(t, todo.ID, id)
	assert.Equal(t, todo.Content, content)
}

func TestTodoDeleteById(t *testing.T) {
	conn := dao.ConnectDB()
	todo_repository := &repository.TodoRepository{Conn: conn}
	todos, _ := todo_repository.FindAll()
	id := todos[len(todos)-1].ID
	err := todo_repository.DeleteById(id)

	assert.Equal(t, err, nil)
}
