package repository

import (
	"app/cmd/domain/model"

	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	Conn *gorm.DB
}

func (t *TodoRepository) FindAll() ([]model.Todo, error) {
	todos := []model.Todo{}
	err := t.Conn.Find(&todos).Error
	return todos, err
}

func (t *TodoRepository) FindById(id int) (model.Todo, error) {
	todo := model.Todo{}
	err := t.Conn.Find(&todo, id).Error
	return todo, err
}

func (t *TodoRepository) Save(content string) (model.Todo, error) {
	todo := model.Todo{Content: content}
	err := t.Conn.Create(&todo).Error
	return todo, err
}

func (t *TodoRepository) UpdateById(id int, content string) (model.Todo, error) {
	todo := model.Todo{ID: id, Content: content}
	err := t.Conn.Model(&todo).Update("content", content).Error
	return todo, err
}

func (t *TodoRepository) DeleteById(id int) error {
	err := t.Conn.Delete(&model.Todo{}, id).Error
	return err
}
