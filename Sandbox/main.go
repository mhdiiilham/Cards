package main

import (
	"fmt"
)

type social struct {
	twitter string
	instagram string
}

type contact struct {
	email string
	phone int
	socialID social
}

type person struct {
	firstName string
	lastName string
	contact
}

func main() {
	nando := person{
		firstName: "Nando",
		lastName: "Nyata Pasti",
		contact: contact{
			email: "nandoperjaka@milf.com",
			phone: 69,
			socialID: social{
				twitter: "He does not have",
				instagram: "pewyoshino",
			},
		},
	}
	nando.fullName()
	nando.updateLastName("Pasti")
	nando.fullName()
}

func (p person) fullName() {
	fmt.Println(p.firstName+" "+p.lastName)
}

func (p *person) updateLastName(newLastName string) {
	(*p).lastName = newLastName
	// fmt.Printf("%+v", p)
}