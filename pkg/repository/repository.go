package repository

import (
	root "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user root.User) (int, error)
	GetUser(username, password string) (root.User, error)
}

type TodoList interface {
	Create(userId int, list root.TodoList) (int, error)
	GetAllList(userId int) ([]root.TodoList, error)
	GetByIDList(userId, listId int) (root.TodoList, error)
	DeleteByID(userId, listId int) error
	UpdatedByID(userId, listId int, input root.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item root.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]root.TodoItem, error)
	GetByIdItem(userId, itemId int) (root.TodoItem, error)
	DeleteByID(userId, itemId int) error
	UpdatedByID(userId, listId int, input root.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
