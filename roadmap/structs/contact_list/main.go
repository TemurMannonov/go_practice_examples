package main

import (
	"fmt"
)

type Contact struct {
	Id int
	Name string
	Gender string
	Phone int
	Mail string
}

type ContactManager struct {
	contacts []Contact
}

func NewContactManager() ContactManager {
	cm := ContactManager{}
	cm.contacts = []Contact{}

	return cm
}

func (cm *ContactManager) Add(c *Contact) *Contact{
	cm.contacts = append(cm.contacts, *c)
	id := len(cm.contacts) - 1
	return &cm.contacts[id]
}

func (cm *ContactManager) Update(c *Contact) *Contact {
	contact := &cm.contacts[c.Id]
	contact.Name = c.Name
	contact.Gender = c.Gender
	contact.Phone = c.Phone
	contact.Mail = c.Mail
	return contact
}

func (cm *ContactManager) Delete(id int) {
	cm.contacts = append(cm.contacts[:id], cm.contacts[id+1:]...)
}

func (cm *ContactManager) ContactList() {
	for _, c := range cm.contacts {
		c.ContactDetail()
	}
}

func (c *Contact) ContactDetail() {
	fmt.Println("**************************")
	fmt.Println("name:", c.Name)
	fmt.Println("id:", c.Id)
	fmt.Println("gender:", c.Gender)
	fmt.Println("phone:", c.Phone)
	fmt.Println("mail:", c.Mail)
	fmt.Println()
}

func (c *ContactManager) GetAllContacts() []Contact{
	return c.contacts
}


// func Menu() {
// 	fmt.Println("**************************")
// 	fmt.Println("*          Menu          *")
// 	fmt.Println("**************************")

// 	fmt.Println("List of all contacts - 1")
// 	fmt.Println("Add new contact      - 2")
// 	fmt.Println("Update a contact     - 3")
// 	fmt.Println("Delete a contact     - 4")
// 	fmt.Println("Exit                 - 5")
// 	fmt.Println("**************************")
// }

// func EnterDetails(c *Contact){
// 	var phone int
// 	var name, gender, mail string

// 	fmt.Print("Enter name: ")
// 	fmt.Scanln(&name)
// 	fmt.Print("Enter gender: ")
// 	fmt.Scanln(&gender)
// 	fmt.Print("Enter phone: ")
// 	fmt.Scanln(&phone)
// 	fmt.Print("Enter mail: ")
// 	fmt.Scanln(&mail)
	
// 	c.name = name
// 	c.gender = gender
// 	c.phone = phone
// 	c.mail = mail
// }
func main() {

	// var choice int
	// var id int
	// var cm ContactManager
	// var c Contact

	// for {
	// 	Menu()
	// 	fmt.Print("Enter a your choice: ")
	// 	fmt.Scan(&choice)
		
	// 	if choice == 1 {
	// 		cm.ContactList()
	// 	} else if choice == 2 {
	// 		c.id = len(cm.contacts)
	// 		EnterDetails(&c)
	// 		cm.Add(c)
	// 	} else if choice == 3 {
	// 		fmt.Print("Enter id: ")
	// 		fmt.Scanln(&id)

	// 		if id <= len(cm.contacts) - 1 {
	// 			c.id = id
	// 			EnterDetails(&c)
	// 			cm.Update(c)
	// 		} else {
	// 			fmt.Println("Entered id does not exists.")
	// 		}
	// 	} else if choice == 4 {
	// 		fmt.Print("Enter id: ")
	// 		fmt.Scanln(&id)

	// 		if id <= len(cm.contacts) - 1 {
	// 			cm.Delete(id)
	// 		} else {
	// 			fmt.Println("Entered id does not exists.")
	// 		}
	// 	} else if choice == 5 {
	// 		break
	// 	}
	// }
}