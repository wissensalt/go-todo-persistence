package repository

import (
	"database/sql"
)

type (
	Todo struct {
		Id        int
		Task      string
		Completed bool
	}

	TodoRepository interface {
		Create(todo Todo) error
		Update(todo Todo) error
		Get() []Todo
		GetById(id int) Todo
		Delete(id int) error
	}

	TodoRepositoryImpl struct {
		*sql.DB
	}
)

func (t TodoRepositoryImpl) Create(todo Todo) error {
	_, err := t.Exec("INSERT INTO todo(task, completed) VALUES($1, $2)", todo.Task, todo.Completed)

	return err
}

func (t TodoRepositoryImpl) Update(todo Todo) error {
	_, err := t.Exec("UPDATE todo SET task=$1, completed=$2 WHERE id=$3", todo.Task, todo.Completed, todo.Id)

	return err
}

func (t TodoRepositoryImpl) Get() []Todo {
	rows, err := t.Query("SELECT * FROM todo")
	if err != nil {
		return []Todo{}
	}
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Task, &todo.Completed)
		if err != nil {
			return nil
		}
		todos = append(todos, todo)
	}

	return todos
}

func (t TodoRepositoryImpl) GetById(id int) Todo {
	r := t.QueryRow("SELECT * FROM todo WHERE id=$1", id)
	var todo Todo
	err := r.Scan(&todo.Id, &todo.Task, &todo.Completed)
	if err != nil {
		return Todo{}
	}

	return todo
}

func (t TodoRepositoryImpl) Delete(id int) error {
	_, err := t.Exec("DELETE FROM todo WHERE id=$1", id)
	return err
}
