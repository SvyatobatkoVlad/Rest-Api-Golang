package repository

import (
	"fmt"
	root "github.com/SvyatobatkoVlad/Rest-Api-Golang"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (t *TodoItemPostgres) Create(listId int, item root.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (t *TodoItemPostgres) GetAllItems(userId, listId int) ([]root.TodoItem, error) {
	var items []root.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id 
    INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemTable, listsItemsTable, usersListsTable)
	if err := t.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (t *TodoItemPostgres) GetByIdItem(userId, itemId int) (root.TodoItem, error) {
	var item root.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
	INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemTable, listsItemsTable, usersListsTable)
	err := t.db.Get(&item, query, itemId, userId)

	return item, err
}

func (t *TodoItemPostgres) DeleteByID(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s tl USING %s li, %s ul 
       WHERE tl.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemTable, listsItemsTable, usersListsTable)
	_, err := t.db.Exec(query, userId, itemId)
	return err
}

func (t *TodoItemPostgres) UpdatedByID(userId, itemId int, input root.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
										WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	logrus.Debugf("updateQuery: %s", query)
	args = append(args, userId, itemId)

	_, err := t.db.Exec(query, args...)
	return err
}
