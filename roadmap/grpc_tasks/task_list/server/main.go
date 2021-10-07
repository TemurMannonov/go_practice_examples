package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"task_list/models"
	"task_list/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	tm models.TaskManagerI
}

const (
	username = "temur"
	password = "12345"
	dbname   = "test"
	sslmode  = "disable"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)
	tm, err := models.NewTaskManager(conn)

	if err != nil {
		log.Fatal("Can not connecting to database ...", err)
	}

	srv := grpc.NewServer()
	proto.RegisterTaskManagerServer(srv, &server{tm})
	reflection.Register(srv)
	fmt.Println("Server is running on port 4040 ...")
	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, t *proto.Task) (*proto.Task, error) {
	err := s.tm.Add(t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *server) Update(ctx context.Context, t *proto.Task) (*proto.Task, error) {
	err := s.tm.Update(t.GetId(), t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *server) MakeDone(ctx context.Context, t *proto.TaskId) (*proto.Status, error) {
	err := s.tm.MakeDone(t.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.Status{Status: "success"}, nil
}

func (s *server) Delete(ctx context.Context, t *proto.TaskId) (*proto.Status, error) {
	err := s.tm.Delete(t.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.Status{Status: "success"}, nil
}

func (s *server) GetAll(ctx context.Context, e *empty.Empty) (*proto.Tasks, error) {
	tasks, err := s.tm.GetAll()
	if err != nil {
		return nil, err
	}

	return &proto.Tasks{Tasks: tasks}, nil
}

func (s *server) GetFinished(ctx context.Context, e *empty.Empty) (*proto.Tasks, error) {
	tasks, err := s.tm.GetFinished()
	if err != nil {
		return nil, err
	}

	return &proto.Tasks{Tasks: tasks}, nil
}

func (s *server) GetNotFinished(ctx context.Context, e *empty.Empty) (*proto.Tasks, error) {
	tasks, err := s.tm.GetNotFinished()
	if err != nil {
		return nil, err
	}

	return &proto.Tasks{Tasks: tasks}, nil
}
