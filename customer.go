package main

type Customer struct {
	Name    string
	Surname string
	Age     int
	Money   float64
}

func NewCustomer(name, surname string, age int, money float64) *Customer {
	return &Customer{
		Name:    name,
		Surname: surname,
		Age:     age,
		Money:   money,
	}
}
