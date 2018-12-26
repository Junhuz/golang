package model

import "fmt"

type Model struct {
	ID    uint32 `json:"user_id"`
	Name  string
	Phone int
	Email string
}

type Customer struct {
	Name  string
	Phone int
}

func (c Customer) GetName() string {
	return c.Name
}

func (c Customer) SetName(name string) {
	c.Name = name
	fmt.Println(c.Name)
}
