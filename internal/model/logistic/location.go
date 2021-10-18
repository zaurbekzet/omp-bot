package logistic

import "fmt"

type Location struct {
	ID        uint64
	Latitude  float64
	Longitude float64
	Title     string
}

func (l Location) String() string {
	return fmt.Sprintf("Location %d: %s (%f, %f)", l.ID, l.Title, l.Latitude, l.Longitude)
}
