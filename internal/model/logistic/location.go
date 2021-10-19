package logistic

import "fmt"

type Location struct {
	ID        uint64  `json:"id,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
}

func (l Location) String() string {
	return fmt.Sprintf("Location %d: %s (%f, %f)", l.ID, l.Title, l.Latitude, l.Longitude)
}
