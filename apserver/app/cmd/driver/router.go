package driver

import (
	presentaion "app/cmd/interface/presentation"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router(conn *gorm.DB) *gin.Engine {
	engine := gin.Default()
	todo := presentaion.InteractorTodo(conn)

	engine.GET("/", presentaion.HealthCheck)
	engine.GET("/api/todos", todo.Index)
	engine.GET("/api/todos/:id", todo.Show)
	engine.POST("/api/todos", todo.Create)
	engine.PUT("/api/todos/:id", todo.Update)
	engine.DELETE("/api/todos/:id", todo.Delete)

	return engine
}
