package models

import (
	_ "database/sql"
	"fmt"

	pr "contact_list/proto"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ContactManager struct {
	db *sqlx.DB
}

func NewContactManager(db string) (ContactManagerI, error) {
	cm := ContactManager{}
	var err error
	cm.db, err = sqlx.Connect("postgres", db)

	if err != nil {
		return nil, err
	}
	return &cm, nil
}

func (cm *ContactManager) Add(c *pr.Contact) error {
	query := "INSERT INTO contact(name, age, gender, phone) VALUES($1, $2, $3, $4)"
	cm.db.MustExec(query, c.Name, c.Age, c.Gender, c.Phone)

	return nil
}

func (cm *ContactManager) Update(id int64, c *pr.Contact) error {
	query := "UPDATE contact SET name=$1, age=$2, gender=$3, phone=$4 WHERE id=$5"
	cm.db.MustExec(query, c.Name, c.Age, c.Gender, c.Phone, id)

	return nil
}

func (cm *ContactManager) Delete(id int64) error {
	query := "DELETE FROM contact WHERE id=$1"
	cm.db.MustExec(query, id)

	return nil
}

func (cm *ContactManager) GetAll() ([]*pr.Contact, error) {
	var contacts []*pr.Contact

	listAllQuery := "select * from contact"

	err := cm.db.Select(&contacts, listAllQuery)

	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (cm *ContactManager) ListAll() error {
	var contacts []*pr.Contact

	listAllQuery := "select * from contact"

	err := cm.db.Select(&contacts, listAllQuery)

	if err != nil {
		return err
	}
	fmt.Println(contacts)
	return nil
}
