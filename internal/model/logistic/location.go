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

func (l Location) Validate() error {
	if l.Title == "" {
		return fmt.Errorf("location title must not be empty")
	} else if l.Latitude < -90 || l.Latitude > 90 {
		return fmt.Errorf("latitude %f is not in range [-90, 90]", l.Latitude)
	} else if l.Longitude < -180 || l.Longitude > 180 {
		return fmt.Errorf("longitude %f is not in range [-180, 180]", l.Longitude)
	}
	return nil
}
