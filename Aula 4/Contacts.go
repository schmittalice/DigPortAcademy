package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

func main() {

	contacts := make(map[int]Contact)

	contacts[0] = Contact{
		Name:  "Allan",
		Phone: "999-222-333",
	}

	contacts[1] = Contact{
		Name:  "Denis",
		Phone: "888-222-333",
	}

	contacts[2] = Contact{
		Name:  "Iloiva",
		Phone: "777-222-333",
	}

	for _, contact := range contacts {
		fmt.Println("Name:", contact.Name)
		fmt.Println("Phone:", contact.Phone)
		fmt.Println()
	}
}
