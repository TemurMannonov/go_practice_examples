package models

import (
	pr "contact_list/proto"
)

type ContactManagerI interface {
	Add(c *pr.Contact) error
	Update(id int64, c *pr.Contact) error
	Delete(id int64) error
	GetAll() ([]*pr.Contact, error)
	ListAll() error
}
