package main

import (
	"fmt"
	"bufio"
	"os"
)

type Task struct {
	id int
	assignee string
	title string
	deadline string
	done bool
}

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) Add(t Task) {
	tm.tasks = append(tm.tasks, t)
}

func (tm *TaskManager) Update(t Task) {
	task := &tm.tasks[t.id]
	task.assignee = t.assignee
	task.title = t.title
	task.deadline = t.deadline
}

func (tm *TaskManager) MakeDone(id int) {
	tm.tasks[id].done = true
}

func (tm *TaskManager) Delete(id int) {
	tm.tasks = append(tm.tasks[:id], tm.tasks[id+1:]...)
}

func (tm *TaskManager) TaskList() {
	for _, t := range tm.tasks {
		t.TaskDetail()
	}
}

func (tm *TaskManager) FinishedTaskList() {
	for _, t := range tm.tasks {
		if t.done {
			t.TaskDetail()
		}
	}
}

func (tm *TaskManager) NotFinishedTaskList() {
	for _, t := range tm.tasks {
		if !t.done {
			t.TaskDetail()
		}
	}
}

func (t *Task) TaskDetail() {
	fmt.Println()
	fmt.Println("Id:", t.id)
	fmt.Println("Assignee:", t.assignee)
	fmt.Println("Title:", t.title)
	fmt.Println("Deadline:", t.deadline)
	fmt.Println("Done:", t.done)
	fmt.Println("--------------------------")
}


func Menu() {
	fmt.Println("**************************")
	fmt.Println("*          Menu          *")
	fmt.Println("**************************")
	fmt.Println("List of all tasks      - 1")
	fmt.Println("Finished tasks         - 2")
	fmt.Println("Not finished tasks     - 3")
	fmt.Println("Add new task           - 4")
	fmt.Println("Update a task          - 5")
	fmt.Println("Make a task done       - 6")
	fmt.Println("Delete a task          - 7")
	fmt.Println("Exit                   - 8")
	fmt.Println("**************************")
}

func EnterDetails(t *Task) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter assignee: ")
	scanner.Scan()
	t.assignee = scanner.Text()

	fmt.Print("Enter title: ")
	scanner.Scan()
	t.title = scanner.Text()

	fmt.Print("Enter deadline1(y-m-d): ")
	scanner.Scan()
	t.deadline = scanner.Text()
}
func main() {

	var choice int
	var id int
	var tm TaskManager
	var t Task

	for {
		Menu()
		fmt.Print("Enter a your choice: ")
		fmt.Scan(&choice)
		
		if choice == 1 {
			tm.TaskList()
		} else if choice == 2 {
			tm.FinishedTaskList()
		} else if choice == 3 {
			tm.NotFinishedTaskList()
		} else if choice == 4 {
			t.id = len(tm.tasks)
			EnterDetails(&t)
			tm.Add(t)
		} else if choice == 5 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)

			if id <= len(tm.tasks) - 1 {
				t.id = id
				EnterDetails(&t)
				tm.Update(t)
			} else {
				fmt.Println("Entered id does not exists.")
			}
		} else if choice == 6 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)

			if id <= len(tm.tasks) - 1 {
				tm.MakeDone(id)
			} else {
				fmt.Println("Entered id does not exists.")
			}
		} else if choice == 7 {
			fmt.Print("Enter id: ")
			fmt.Scanln(&id)

			if id <= len(tm.tasks) - 1 {
				tm.Delete(id)
			} else {
				fmt.Println("Entered id does not exists.")
			}
		} else if choice == 8 {
			break
		}
	}
}