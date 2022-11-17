package service

import (
	root "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list root.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAllList(userId int) ([]root.TodoList, error) {
	return s.repo.GetAllList(userId)
}

func (s *TodoListService) GetByIDList(userId, listId int) (root.TodoList, error) {
	return s.repo.GetByIDList(userId, listId)
}
