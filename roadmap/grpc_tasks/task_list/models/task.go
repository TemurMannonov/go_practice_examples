package models

import (
	_ "database/sql"
	"fmt"

	pr "task_list/proto"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TaskManager struct {
	db *sqlx.DB
}

func NewTaskManager(db string) (TaskManagerI, error) {
	tm := TaskManager{}
	var err error
	tm.db, err = sqlx.Connect("postgres", db)

	if err != nil {
		return nil, err
	}
	return &tm, nil
}

func (tm *TaskManager) Add(t *pr.Task) error {
	query := "INSERT INTO tasks_grpc(assignee, title, deadline) VALUES($1, $2, $3)"
	tm.db.MustExec(query, t.Assignee, t.Title, t.Deadline)

	return nil
}

func (tm *TaskManager) Update(id int64, t *pr.Task) error {
	query := "UPDATE tasks_grpc SET assignee=$1, title=$2, deadline=$3 WHERE id=$4"
	tm.db.MustExec(query, t.Assignee, t.Title, t.Deadline, id)

	return nil
}

func (tm *TaskManager) MakeDone(id int64) error {
	query := "UPDATE tasks_grpc SET done=$1 WHERE id=$2"
	tm.db.MustExec(query, true, id)

	return nil
}

func (tm *TaskManager) Delete(id int64) error {
	query := "DELETE FROM tasks_grpc WHERE id=$1"
	tm.db.MustExec(query, id)

	return nil
}

func (tm *TaskManager) GetAll() ([]*pr.Task, error) {
	var task []*pr.Task

	query := "SELECT * FROM tasks_grpc"

	err := tm.db.Select(&task, query)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tm *TaskManager) GetFinished() ([]*pr.Task, error) {
	var task []*pr.Task

	query := "SELECT * FROM tasks_grpc WHERE done=True"

	err := tm.db.Select(&task, query)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tm *TaskManager) GetNotFinished() ([]*pr.Task, error) {
	var task []*pr.Task

	query := "SELECT * FROM tasks_grpc WHERE done=False"

	err := tm.db.Select(&task, query)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tm *TaskManager) ListAll() error {
	var tasks []*pr.Task

	query := "select * from tasks_grpc"

	err := tm.db.Select(&tasks, query)

	if err != nil {
		return err
	}
	fmt.Println(tasks)
	return nil
}
