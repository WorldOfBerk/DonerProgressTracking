package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 10 tane customer ac
	// owner instance ac
	// doner ac
	// storage ac
	// doneri storage icine at
	// storage owner icine at

	doners := map[string]*Doner{
		"ceyrek":    NewDoner("Ceyrek", 3.0, 2.0),
		"yarim":     NewDoner("Yarim", 5.0, 2.0),
		"üc-ceyrek": NewDoner("Üc-Ceyrek", 7.0, 8.0),
		"tam":       NewDoner("Tam", 10.0, 12.0),
	}

	storage := NewStorage(10, 100.0, doners["tam"])

	owner := NewOwner(storage)

	customers := []*Customer{
		NewCustomer("John", "Doe", 30, 27.0),
		NewCustomer("Jane", "Doen", 25, 10.9),
		NewCustomer("Johnny", "Doenny", 18, 12.0),
		NewCustomer("Johann", "Derek", 44, 45.0),
		NewCustomer("Joe", "Dolap", 35, 32.0),
		NewCustomer("Joel", "Deleven", 20, 21.6),
		NewCustomer("Joline", "Daneyo", 33, 20.0),
		NewCustomer("Jeremy", "Dung", 32, 22.0),
		NewCustomer("Jacuan", "Dex", 31, 25.0),
		NewCustomer("Jinx", "Dimitri", 30, 20.5),
	}
	totalSpent := 0.0
	for _, customer := range customers {
		doner := getDoner(doners)
		fmt.Printf("%s %s, (%d yasli), %.2f€ ile %s döner istiyor\n", customer.Name, customer.Surname, customer.Age, customer.Money, doner)

		donertane := doners[doner]
		totalPrice := (donertane.Duration * donertane.Percentage) / 100

		if customer.Money >= totalPrice {
			customer.Money -= totalPrice
			totalSpent += totalPrice
			storage.Rolls--
			storage.Durability -= owner.Storage.Doner.Percentage
			fmt.Printf("%s %s %s doner satin aldi. Kalan Para: %.2f€, Kalan döner: %d, Depo Durablitiy: %.2f\n", customer.Name, customer.Surname, doner, customer.Money, storage.Rolls, storage.Durability)
		} else {
			fmt.Println("No money ")
		}
	}
	fmt.Printf("Kalan doner: %d, toplam para: %.2f€\n", storage.Rolls, totalSpent)
}

func getDoner(doners map[string]*Doner) string {
	types := make([]string, 0, len(doners))
	for doner := range doners {
		types = append(types, doner)
	}
	return types[rand.Intn(len(types))]
}
