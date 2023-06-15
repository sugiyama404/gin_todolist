package todotest

import (
	"app/cmd/interface/repository"
	"app/cmd/utils/testtool"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoIndex(t *testing.T) {
	router, w, _ := testtool.NewTestServerWithRouter()
	req, _ := http.NewRequest("GET", "/api/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestTodoCreate(t *testing.T) {
	router, w, _ := testtool.NewTestServerWithRouter()
	bodyReader := strings.NewReader(`{"content": "test"}`)
	req, _ := http.NewRequest("POST", "/api/todos", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestTodoShow(t *testing.T) {
	router, w, conn := testtool.NewTestServerWithRouter()
	todo_repository := &repository.TodoRepository{Conn: conn}
	todos, _ := todo_repository.FindAll()
	id := strconv.Itoa(todos[len(todos)-1].ID)
	req, _ := http.NewRequest("GET", "/api/todos/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestTodoUpdate(t *testing.T) {
	router, w, conn := testtool.NewTestServerWithRouter()
	todo_repository := &repository.TodoRepository{Conn: conn}
	todos, _ := todo_repository.FindAll()
	id := strconv.Itoa(todos[len(todos)-1].ID)
	bodyReader := strings.NewReader(`{"id": '` + id + `', "content": "test"}`)
	req, _ := http.NewRequest("PUT", "/api/todos/"+id, bodyReader)
	router.ServeHTTP(w, req)
	assert.Equal(t, 202, w.Code)
}

func TestTodoDelete(t *testing.T) {
	router, w, conn := testtool.NewTestServerWithRouter()
	todo_repository := &repository.TodoRepository{Conn: conn}
	todos, _ := todo_repository.FindAll()
	id := strconv.Itoa(todos[len(todos)-1].ID)
	req, _ := http.NewRequest("DELETE", "/api/todos/"+id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 202, w.Code)
}
