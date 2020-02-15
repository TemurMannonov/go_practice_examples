package models

import (
	"contact_list/proto"
	"context"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var client proto.ContactManagerClient

func Test_Add(t *testing.T) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewContactManagerClient(conn)

	c1 := &proto.Contact{
		Name:   "Umar",
		Age:    "42",
		Gender: "Male",
		Phone:  "9983242",
	}

	contact, err := client.Add(context.Background(), c1)
	if err != nil {
		t.Error("Error while adding a contact: ", err)
	}
	fmt.Println("Response:", contact)
}

func Test_Update(t *testing.T) {

	c1 := &proto.Contact{
		Id:     54,
		Name:   "Muhammad",
		Age:    "42",
		Gender: "Male",
		Phone:  "9983242",
	}

	contact, err := client.Update(context.Background(), c1)
	if err != nil {
		t.Error("Error while updating a contact: ", err)
	}
	fmt.Println("Response: ", contact)
}

func TestContactManager_Delete(t *testing.T) {
	id := &proto.ContactId{
		Id: 53,
	}

	status, err := client.Delete(context.Background(), id)
	if err != nil {
		t.Error("Error while deleting a contact: ", err)
	}
	fmt.Println("Response: ", status)
}

func Test_GetAll(t *testing.T) {
	contacts, err := client.GetAll(context.Background(), &empty.Empty{})
	if err != nil {
		t.Error("Error while getting all contacts: ", err)
	}
	fmt.Println("Response: ", contacts)

}
