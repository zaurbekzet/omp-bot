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

var InitialLocations = []Location{
	{
		ID:        1,
		Latitude:  -13.163060,
		Longitude: -72.545560,
		Title:     "Machu Picchu",
	},
	{
		ID:        2,
		Latitude:  20.683060,
		Longitude: -88.568610,
		Title:     "Chichen Itza",
	},
	{
		ID:        3,
		Latitude:  41.890169,
		Longitude: 12.492269,
		Title:     "Colosseum",
	},
	{
		ID:        4,
		Latitude:  27.174931,
		Longitude: 78.042097,
		Title:     "Taj Mahal",
	},
	{
		ID:        5,
		Latitude:  29.979167,
		Longitude: 31.134167,
		Title:     "Great Pyramid of Giza",
	},
}
