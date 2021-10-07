package main

import (
	"fmt"
	// "bufio"
	// "os"
)

type Task struct {
	Id int
	Assignee string
	Title string
	Deadline string
	Done bool
}

type TaskManager struct {
	tasks []Task
}

func NewTaskManager() TaskManager {
	cm := TaskManager{}
	cm.tasks = []Task{}

	return cm
}

func (tm *TaskManager) Add(t *Task) *Task {
	tm.tasks = append(tm.tasks, *t)
	id := len(tm.tasks) - 1
	return &tm.tasks[id]
}

func (tm *TaskManager) Update(t *Task) *Task{
	task := &tm.tasks[t.Id]
	task.Assignee = t.Assignee
	task.Title = t.Title
	task.Deadline = t.Deadline

	return task
}

func (tm *TaskManager) MakeDone(Id int) {
	tm.tasks[Id].Done = true
}

func (tm *TaskManager) Delete(Id int) {
	tm.tasks = append(tm.tasks[:Id], tm.tasks[Id+1:]...)
}

func (tm *TaskManager) TaskList() []Task{
	// for _, t := range tm.tasks {
	// 	t.TaskDetail()
	// }
	return tm.tasks
}

func (tm *TaskManager) FinishedTaskList() []Task{
	tasks := []Task{}
	for _, t := range tm.tasks {
		if t.Done {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (tm *TaskManager) NotFinishedTaskList() []Task {
	tasks := []Task{}
	for _, t := range tm.tasks {
		if !t.Done {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (t *Task) TaskDetail() {
	fmt.Println()
	fmt.Println("Id:", t.Id)
	fmt.Println("Assignee:", t.Assignee)
	fmt.Println("Title:", t.Title)
	fmt.Println("Deadline:", t.Deadline)
	fmt.Println("Done:", t.Done)
	fmt.Println("--------------------------")
}


// func Menu() {
// 	fmt.Println("**************************")
// 	fmt.Println("*          Menu          *")
// 	fmt.Println("**************************")
// 	fmt.Println("List of all tasks      - 1")
// 	fmt.Println("Finished tasks         - 2")
// 	fmt.Println("Not finished tasks     - 3")
// 	fmt.Println("Add new task           - 4")
// 	fmt.Println("Update a task          - 5")
// 	fmt.Println("Make a task Done       - 6")
// 	fmt.Println("Delete a task          - 7")
// 	fmt.Println("Exit                   - 8")
// 	fmt.Println("**************************")
// }

// func EnterDetails(t *Task) {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Print("Enter Assignee: ")
// 	scanner.Scan()
// 	t.Assignee = scanner.Text()

// 	fmt.Print("Enter Title: ")
// 	scanner.Scan()
// 	t.Title = scanner.Text()

// 	fmt.Print("Enter Deadline1(y-m-d): ")
// 	scanner.Scan()
// 	t.Deadline = scanner.Text()
// }
func main() {

	// var choice int
	// var Id int
	// var tm TaskManager
	// var t Task

	// for {
	// 	Menu()
	// 	fmt.Print("Enter a your choice: ")
	// 	fmt.Scan(&choice)
		
	// 	if choice == 1 {
	// 		tm.TaskList()
	// 	} else if choice == 2 {
	// 		tm.FinishedTaskList()
	// 	} else if choice == 3 {
	// 		tm.NotFinishedTaskList()
	// 	} else if choice == 4 {
	// 		t.Id = len(tm.tasks)
	// 		EnterDetails(&t)
	// 		tm.Add(t)
	// 	} else if choice == 5 {
	// 		fmt.Print("Enter Id: ")
	// 		fmt.Scanln(&Id)

	// 		if Id <= len(tm.tasks) - 1 {
	// 			t.Id = Id
	// 			EnterDetails(&t)
	// 			tm.Update(t)
	// 		} else {
	// 			fmt.Println("Entered Id does not exists.")
	// 		}
	// 	} else if choice == 6 {
	// 		fmt.Print("Enter Id: ")
	// 		fmt.Scanln(&Id)

	// 		if Id <= len(tm.tasks) - 1 {
	// 			tm.MakeDone(Id)
	// 		} else {
	// 			fmt.Println("Entered Id does not exists.")
	// 		}
	// 	} else if choice == 7 {
	// 		fmt.Print("Enter Id: ")
	// 		fmt.Scanln(&Id)

	// 		if Id <= len(tm.tasks) - 1 {
	// 			tm.Delete(Id)
	// 		} else {
	// 			fmt.Println("Entered Id does not exists.")
	// 		}
	// 	} else if choice == 8 {
	// 		break
	// 	}
	// }
}