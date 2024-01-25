package main

type Doner struct {
	Name       string
	Duration   float64
	Percentage float64
}

func NewDoner(name string, duration, percentage float64) *Doner {
	return &Doner{
		Name:       name,
		Duration:   duration,
		Percentage: percentage,
	}
}
