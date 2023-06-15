package presentaion

import (
	"app/cmd/domain/form"
	"app/cmd/interface/repository"
	"app/cmd/usecase"
	"app/cmd/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Interactor usecase.UseCase
}

func InteractorTodo(conn *gorm.DB) *Todo {
	return &Todo{
		Interactor: usecase.UseCase{
			Repository: repository.TodoRepository{
				Conn: conn,
			},
		},
	}
}

func (t *Todo) Index(c *gin.Context) {
	todos, err := t.Interactor.GetAll()
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (t *Todo) Show(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	todo, err := t.Interactor.GetById(id)
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (t *Todo) Create(c *gin.Context) {
	var request form.TodoRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	content := request.Content
	todo, err := t.Interactor.Create(content)
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (t *Todo) Update(c *gin.Context) {
	var request form.TodoRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	id_string := c.Param("id")
	content := request.Content
	id, err := strconv.Atoi(id_string)
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	todo, err := t.Interactor.Update(id, content)
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusAccepted, todo)
}

func (t *Todo) Delete(c *gin.Context) {
	id_string := c.Param("id")
	id, err := strconv.Atoi(id_string)
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	err = t.Interactor.Delete(id)
	if err != nil {
		apiErr := errors.BadRequestHandler(err, http.StatusBadRequest)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "ok"})
}
