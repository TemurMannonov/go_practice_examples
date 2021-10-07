package main

import (
	"testing"
)

var tm = NewTaskManager()

func TestTaskManager_Add(t *testing.T) {

	task := &Task{Assignee: "Rustam", Title: "Backend Developer", Deadline: "2020-09-09"}
	addedTask := tm.Add(task)
	if addedTask.Assignee != task.Assignee {
		t.Error("Error while adding task")
	}
}

func TestTaskManager_Update(t *testing.T) {
	task := &tm.tasks[0]
	task.Assignee = "Temur"
	task.Title = "Devops"
	task.Deadline = "2020-02-20"
	updatedTask := tm.Update(task)
	if updatedTask != task {
		t.Error("Error while updating task")
	}
}

func TestTaskManager_List(t *testing.T) {
	tasks := tm.TaskList()
	if len(tasks) != 1 && tasks[0].Assignee != "Temur" {
		t.Error("Error while returning task list")
	}
}

func TestTaskManager_NotFinishedList(t *testing.T) {
	tasks := tm.NotFinishedTaskList()
	if len(tasks) != 1 && tasks[0].Assignee != "Temur" {
		t.Error("Error while returning not finished task list")
	}
}

func TestTaskManager_MakeDone(t *testing.T) {
	tm.MakeDone(0)
	if !tm.tasks[0].Done {
		t.Error("Error while making done task")
	}
}

func TestTaskManager_FinishedList(t *testing.T) {
	tasks := tm.FinishedTaskList()
	if len(tasks) != 1 && tasks[0].Assignee != "Temur" {
		t.Error("Error while returning finished task list")
	}
}

func TestTaskManager_Delete(t *testing.T) {
	tm.Delete(0)
	if len(tm.tasks) != 0 {
		t.Error("Error while deleting task")
	}
}

