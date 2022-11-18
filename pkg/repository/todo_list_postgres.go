package repository

import (
	"fmt"
	root "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (t *TodoListPostgres) Create(userId int, list root.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *TodoListPostgres) GetAllList(userId int) ([]root.TodoList, error) {
	var lists []root.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListTable, usersListsTable)
	err := t.db.Select(&lists, query, userId)

	return lists, err
}

func (t *TodoListPostgres) GetByIDList(userId, listId int) (root.TodoList, error) {
	var lists root.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListTable, usersListsTable)
	err := t.db.Get(&lists, query, userId, listId)

	return lists, err
}

func (t *TodoListPostgres) DeleteByID(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %S ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListTable, usersListsTable)
	_, err := t.db.Exec(query, userId, listId)
	return err
}

func (t *TodoListPostgres) UpdatedByID(userId, listId int, input root.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %S tl SET %s FROM %S ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListTable, setQuery, usersListsTable, argId, argId+1)
	logrus.Debugf("updateQuery: %s", query)

	_, err := t.db.Exec(query, args...)
	return err
}
