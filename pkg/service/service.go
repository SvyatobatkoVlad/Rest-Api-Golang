package service

import (
	root "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user root.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list root.TodoList) (int, error)
	GetAllList(userId int) ([]root.TodoList, error)
	GetByIDList(userId, listId int) (root.TodoList, error)
	DeleteByID(userId, listId int) error
	UpdatedByID(userId, listId int, input root.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		TodoList:      NewTodoListService(repo),
	}
}
