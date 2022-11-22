package service

import (
	root "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"github.com/SvyatobatkoVlad/Rest-Api-Golang/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item root.TodoItem) (int, error) {
	_, err := s.listRepo.GetByIDList(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAllItems(userId, listId int) ([]root.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetByIdItem(userId, itemId int) (root.TodoItem, error) {
	return s.repo.GetByIdItem(userId, itemId)
}

func (s *TodoItemService) DeleteByID(userId, itemId int) error {
	return s.repo.DeleteByID(userId, itemId)
}

func (s *TodoItemService) UpdatedByID(userId, listId int, input root.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdatedByID(userId, listId, input)
}
