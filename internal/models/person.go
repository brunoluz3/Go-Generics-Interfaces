package models

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s  %s, idade: %d", p.FirstName, p.LastName, p.Age)
}
