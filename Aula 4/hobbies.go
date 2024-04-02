package main

import "fmt"

type Contact struct {
	Name   string
	Phone  string
	Email  string
	Hobbie string
}

func main() {
	contacts := make(map[int]Contact)

	contacts[0] = Contact{
		Name:   "John Doe",
		Phone:  "123-456-7890",
		Email:  "johndoe@example.com",
		Hobbie: "to read",
	}

	contacts[1] = Contact{
		Name:   "Jane Smith",
		Phone:  "098-765-4321",
		Email:  "janesmith@example.com",
		Hobbie: "to play soccer",
	}

	contacts[2] = Contact{
		Name:   "Bob Johnson",
		Phone:  "111-222-3333",
		Email:  "bobjohnson@example.com",
		Hobbie: "to watch tv",
	}

	for _, contact := range contacts {
		fmt.Println("Name:", contact.Name)
		fmt.Println("Phone:", contact.Phone)
		fmt.Println("Email:", contact.Email)
		fmt.Println("Hobbie:", contact.Hobbie)
		fmt.Println()
	}
}
