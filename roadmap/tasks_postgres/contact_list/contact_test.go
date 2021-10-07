package models

import (
	"testing"
)

var cm ContactManagerI
var id int

func TestNewContactManager(t *testing.T) {
	var err error
	cm, err = NewContactManager()
	if err != nil {
		t.Error("Can not connect a database: ", err)
	}
}
func TestContactManager_Add(t *testing.T) {
	c := Contact{Name: "Rustam", Age: 24, Gender: "male", Phone: "998998149233"}
	err := cm.Add(c)
	if err != nil {
		t.Error("Can not create a contact: ", err)
	}
}

func TestContactManager_Update(t *testing.T) {
	c := Contact{Name: "Temur", Age: 22, Gender: "male", Phone: "998998149233"}

	err := cm.Update(1, c)
	if err != nil {
		t.Error("Can not update a contact: ", err)
	}
}

func TestContactManager_GetAll(t *testing.T) {
	_, err := cm.GetAll()
	if err != nil {
		t.Error("Can not retrieve contacts: ", err)
	}

}

func TestContactManager_ListAll(t *testing.T) {
	err := cm.ListAll()
	if err != nil {
		t.Error("Can not display contacts: ", err)
	}
}

func TestContactManager_Delete(t *testing.T) {
	err := cm.Delete(50)
	if err != nil {
		t.Error("Can not delete contact: ", err)
	}
}
