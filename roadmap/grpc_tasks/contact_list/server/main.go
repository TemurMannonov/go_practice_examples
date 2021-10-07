package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"contact_list/models"
	"contact_list/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	cm models.ContactManagerI
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
	cm, err := models.NewContactManager(conn)

	if err != nil {
		log.Fatal("Can not connecting to database ...", err)
	}

	srv := grpc.NewServer()
	proto.RegisterContactManagerServer(srv, &server{cm})
	reflection.Register(srv)
	fmt.Println("Server is running on port 4040 ...")
	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, c *proto.Contact) (*proto.Contact, error) {
	err := s.cm.Add(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *server) Update(ctx context.Context, c *proto.Contact) (*proto.Contact, error) {
	err := s.cm.Update(c.GetId(), c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *server) Delete(ctx context.Context, c *proto.ContactId) (*proto.Status, error) {
	err := s.cm.Delete(c.GetId())
	if err != nil {
		return nil, err
	}

	return &proto.Status{Status: "success"}, nil
}

func (s *server) GetAll(ctx context.Context, e *empty.Empty) (*proto.GetContacts, error) {
	contacts, err := s.cm.GetAll()
	if err != nil {
		return nil, err
	}

	return &proto.GetContacts{Contacts: contacts}, nil
}
