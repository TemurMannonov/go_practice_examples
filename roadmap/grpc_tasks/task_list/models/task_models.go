package models

import (
	pr "task_list/proto"
)

type TaskManagerI interface {
	Add(c *pr.Task) error
	Update(id int64, c *pr.Task) error
	MakeDone(id int64) error
	Delete(id int64) error
	GetAll() ([]*pr.Task, error)
	GetFinished() ([]*pr.Task, error)
	GetNotFinished() ([]*pr.Task, error)
	ListAll() error
}
