package models

type Contact struct {
	Id     int    `db: "id"`
	Name   string `db: "name"`
	Age    int    `db: "age"`
	Gender string `db: "gender"`
	Phone  string `db: "phone"`
}

type ContactManagerI interface {
	Add(c Contact) error
	Update(id int, c Contact) error
	Delete(id int) error
	GetAll() ([]Contact, error)
	ListAll() error
}
