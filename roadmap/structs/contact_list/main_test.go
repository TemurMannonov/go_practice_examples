package main

import (
	"testing"
)

var cm = NewContactManager()

func TestContactManager_Add(t *testing.T) {

	c := &Contact{Id: 0,Name: "Rustam", Gender: "male", Phone: 998998149233, Mail: "rustam@gmail.com"}
	addedContact := cm.Add(c)
	if addedContact.Name != c.Name {
		t.Error("Error while adding contact")
	}
}

func TestContactManager_Update(t *testing.T) {
	c := &cm.contacts[0]
	c.Name = "Temur"
	c.Gender = "Male"
	c.Phone = 998903599940
	c.Mail = "t.man@gmail.com"
	updatedContact := cm.Update(c)
	if updatedContact != c {
		t.Error("Error while updating contact")
	}
}

func TestContactManager_List(t *testing.T) {
	contacts := cm.GetAllContacts()
	if len(contacts) != 1 && contacts[0].Name != "Temur" {
		t.Error("Error while contact")
	}
}

func TestContactManager_Delete(t *testing.T) {
	cm.Delete(0)
	if len(cm.contacts) != 0 {
		t.Error("Error while deleting contact")
	}
}

