package main

import "fmt"

// User is a representation of a user
type User struct {
	FirstName  string
	LastName   string
	Email string
}
func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// Address is a representation of a Address
type Address struct {
	City    string
	State   string
	Country string
}
func (a Address) Info() string {
	return fmt.Sprintf("%s, %s, %s", a.City, a.State, a.Country)
}

// Customer is an struct representing a customer
type Customer struct {
	// User represents a user
	User
	// Addr represents a location of a customer
	Addr Address
}

func main() {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email: "johndoe@gmail.com",
	}

	addr := Address{
		City:    "Capital Federal",
		State:   "Buenos Aires",
		Country: "Argentina",
	}

	customer := Customer{
		User: user,
		Addr: addr,
	}

	fmt.Println("Fullname:", customer.FullName())
	fmt.Println("Location:", customer.Addr.Info())
}