package main

type Storage struct {
	Rolls      int
	Durability float64
	Doner      *Doner
}

func NewStorage(rolls int, durability float64, doner *Doner) *Storage {
	return &Storage{
		Rolls:      rolls,
		Durability: durability,
		Doner:      doner,
	}
}
