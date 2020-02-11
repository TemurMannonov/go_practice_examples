package models

import (
	_ "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TaskManager struct {
	db *sqlx.DB
}

func NewTaskManager() (TaskManagerI, error) {
	tm := TaskManager{}
	var err error
	tm.db, err = sqlx.Connect("postgres", "user=temur password=12345 dbname=test sslmode=disable")

	if err != nil {
		return nil, err
	}
	return &tm, nil
}

func (tm *TaskManager) Add(t Task) error {
	query := "INSERT INTO task(assignee, title, deadline) VALUES($1, $2, $3)"
	tm.db.MustExec(query, t.Assignee, t.Title, t.Deadline)

	return nil
}

func (tm *TaskManager) Update(id int, t Task) error {
	query := "UPDATE task SET assignee=$1, title=$2, deadline=$3 WHERE id=$4"
	tm.db.MustExec(query, t.Assignee, t.Title, t.Deadline, id)

	return nil
}

func (tm *TaskManager) MakeDone(id int) error {
	query := "UPDATE task SET done=$1 WHERE id=$2"
	tm.db.MustExec(query, true, id)

	return nil
}

func (tm *TaskManager) Delete(id int) error {
	query := "DELETE FROM task WHERE id=$1"
	tm.db.MustExec(query, id)

	return nil
}

func (tm *TaskManager) GetAll() ([]Task, error) {
	var tasks []Task
	rows, err := tm.db.Queryx("Select id, assignee, title, deadline FROM task")

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Assignee, &t.Title, &t.Deadline)
		tasks = append(tasks, t)
	}
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tm *TaskManager) GetFinished() ([]Task, error) {
	var tasks []Task
	rows, err := tm.db.Queryx("Select id, assignee, title, deadline FROM task WHERE done=false")

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Assignee, &t.Title, &t.Deadline)
		tasks = append(tasks, t)
	}
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tm *TaskManager) GetNotFinished() ([]Task, error) {
	var tasks []Task
	rows, err := tm.db.Queryx("Select id, assignee, title, deadline FROM task WHERE done=true")

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Assignee, &t.Title, &t.Deadline)
		tasks = append(tasks, t)
	}
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tm *TaskManager) ListAll() error {
	var tasks []Task
	rows, err := tm.db.Queryx("Select id, assignee, title, deadline FROM task")

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Assignee, &t.Title, &t.Deadline)
		tasks = append(tasks, t)
	}
	if err != nil {
		return err
	}
	fmt.Println(tasks)
	return nil
}
