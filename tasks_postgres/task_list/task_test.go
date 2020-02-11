package models

import (
	"testing"
	"time"
)

var tm TaskManagerI

func TestNewTaskManager(t *testing.T) {
	var err error
	tm, err = NewTaskManager()
	if err != nil {
		t.Error("Can not connect a database: ", err)
	}
}
func TestTaskManager_Add(t *testing.T) {
	time, _ := time.Parse("2006-02-01", "2020-02-02")
	task := Task{Assignee: "Rustam", Title: "Create a api", Deadline: time}
	err := tm.Add(task)
	if err != nil {
		t.Error("Can not create a Task: ", err)
	}
}

func TestTaskManager_Update(t *testing.T) {
	time, _ := time.Parse("2006-02-01", "2020-02-02")
	task := Task{Assignee: "Temur", Title: "Support backend", Deadline: time}

	err := tm.Update(8, task)
	if err != nil {
		t.Error("Can not update a Task: ", err)
	}
}

func TestTaskManager_MakeDone(t *testing.T) {
	err := tm.MakeDone(1)
	if err != nil {
		t.Error("Can not make done a task: ", err)
	}
}

func TestTaskManager_GetAll(t *testing.T) {
	_, err := tm.GetAll()
	if err != nil {
		t.Error("Can not retrieve Tasks: ", err)
	}

}

func TestTaskManager_Finished(t *testing.T) {
	_, err := tm.GetFinished()
	if err != nil {
		t.Error("Can not retrieve Tasks: ", err)
	}
}

func TestTaskManager_NotFinished(t *testing.T) {
	_, err := tm.GetNotFinished()
	if err != nil {
		t.Error("Can not retrieve Tasks: ", err)
	}
}

func TestTaskManager_ListAll(t *testing.T) {
	err := tm.ListAll()
	if err != nil {
		t.Error("Can not display Tasks: ", err)
	}
}

func TestTaskManager_Delete(t *testing.T) {
	err := tm.Delete(12)
	if err != nil {
		t.Error("Can not delete Task: ", err)
	}
}
