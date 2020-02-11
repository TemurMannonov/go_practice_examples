package models

import "time"

type Task struct {
	Id       int       `db: "id"`
	Assignee string    `db: "assignee"`
	Title    string    `db: "title"`
	Deadline time.Time `db: "deadline"`
	Done     bool      `db: "done"`
}

type TaskManagerI interface {
	Add(t Task) error
	Update(id int, t Task) error
	MakeDone(id int) error
	Delete(id int) error
	GetAll() ([]Task, error)
	GetFinished() ([]Task, error)
	GetNotFinished() ([]Task, error)
	ListAll() error
}
