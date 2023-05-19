package aggregate

import (
	"errors"
	"go-ddd-architecture/entity"
	"go-ddd-architecture/valueobject"

	"github.com/google/uuid"
)

var (
	errInvalidPerson = errors.New("a customer has to have name!")
)

type Customer struct {
	// Person is the root entity of customer
	// Which means person.ID is the main identifier of customer

	person       *entity.Person
	products     []entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new customer aggregate
// it will validate that the name is not empty

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, errInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
