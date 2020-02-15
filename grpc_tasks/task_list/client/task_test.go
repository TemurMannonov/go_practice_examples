package models

import (
	"context"
	"fmt"
	"task_list/proto"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var client proto.TaskManagerClient

func Test_Add(t *testing.T) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewTaskManagerClient(conn)

	task := &proto.Task{
		Assignee: "Bobur",
		Title:    "Create a api",
		Deadline: "2020-09-09",
	}

	response, err := client.Add(context.Background(), task)
	if err != nil {
		t.Error("Error while adding a task: ", err)
	}
	fmt.Println("Response:", response)
}

func Test_Update(t *testing.T) {

	task := &proto.Task{
		Id:       12,
		Assignee: "Bobur",
		Title:    "Create a api",
		Deadline: "2020-09-09",
	}

	response, err := client.Update(context.Background(), task)
	if err != nil {
		t.Error("Error while updating a task: ", err)
	}
	fmt.Println("Response: ", response)
}

func TestContactManager_Delete(t *testing.T) {
	id := &proto.TaskId{
		Id: 53,
	}

	status, err := client.Delete(context.Background(), id)
	if err != nil {
		t.Error("Error while deleting a task: ", err)
	}
	fmt.Println("Response: ", status)
}

func Test_GetAll(t *testing.T) {
	response, err := client.GetAll(context.Background(), &empty.Empty{})
	if err != nil {
		t.Error("Error while getting all tasks: ", err)
	}
	fmt.Println("Response: ", response)
}

func Test_GetFinished(t *testing.T) {
	response, err := client.GetFinished(context.Background(), &empty.Empty{})
	if err != nil {
		t.Error("Error while getting all tasks: ", err)
	}
	fmt.Println("Response: ", response)
}

func Test_GetNotFinished(t *testing.T) {
	response, err := client.GetNotFinished(context.Background(), &empty.Empty{})
	if err != nil {
		t.Error("Error while getting all tasks: ", err)
	}
	fmt.Println("Response: ", response)
}
