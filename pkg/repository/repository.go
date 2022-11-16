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
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
