package usecase

import (
	"app/cmd/domain/model"
	"app/cmd/interface/repository"
	"fmt"
)

type UseCase struct {
	Repository repository.TodoRepository
}

func (u *UseCase) GetAll() ([]model.Todo, error) {
	return u.Repository.FindAll()
}

func (u *UseCase) GetById(id int) (model.Todo, error) {
	return u.Repository.FindById(id)
}

func (u *UseCase) Create(content string) (model.Todo, error) {
	fmt.Println(content)
	return u.Repository.Save(content)
}

func (u *UseCase) Update(id int, content string) (model.Todo, error) {
	return u.Repository.UpdateById(id, content)
}

func (u *UseCase) Delete(id int) error {
	return u.Repository.DeleteById(id)
}
