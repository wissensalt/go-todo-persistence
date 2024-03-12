package service

import "github.com/wissensalt/go-todo-persistence/internal/repository"

type (
	TodoService interface {
		Create(todo repository.Todo) error
		Update(todo repository.Todo) error
		Get() []repository.Todo
		GetById(id int) repository.Todo
		Delete(id int) error
	}
	TodoServiceImpl struct {
		repository.TodoRepository
	}
)
