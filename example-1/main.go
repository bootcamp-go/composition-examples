package main

import "fmt"

// ______________________________________________________________
// User is a representation of a user
type User struct {
	Name  string
	Email string
}
func (u User) UserInfo() string {
	return u.Name + " " + u.Email
}

// ______________________________________________________________
// Address is a representation of a Address
type Address struct {
	City    string
	State   string
	Country string
}
func (a Address) AddressInfo() string {
	return fmt.Sprintf("%s, %s, %s", a.City, a.State, a.Country)
}

// ______________________________________________________________
// Customer is a representation of a Customer
type Customer struct {
	User			// Embedding
	Address Address // Embedding
	CUI     string
}

func (c Customer) UserInfo() string {
	msg := "Customer: " + c.User.UserInfo()
	return msg
}

func (c Customer) Info() string {
	return c.User.UserInfo() + " " + c.Address.AddressInfo()
}

// ______________________________________________________________
type Order struct {
	Customer
	OrderID int
	Address
}

type PersonAttributes struct {
	Name string
	Age int
	Address Address
}
type Person struct {
	// Id int
	Id int
	// Attributes PersonAttributes
	PersonAttributes
}


func main() {
	ct := Customer{
		User: User{
			Name:  "John Doe",
			Email: "johndoe@hotmail.com",
		},
		Address: Address{
			City:    "New York",
			State:   "New York",
			Country: "USA",
		},
	}

	// fmt.Println(ct.Info())
	fmt.Println(ct.UserInfo())
}