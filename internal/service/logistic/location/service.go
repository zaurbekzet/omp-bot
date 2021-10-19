package location

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

type LocationService interface {
	Describe(locationID uint64) (*logistic.Location, error)
	List(cursor uint64, limit uint64) ([]logistic.Location, error)
	Create(logistic.Location) (uint64, error)
	Update(locationID uint64, location logistic.Location) error
	Remove(locationID uint64) (bool, error)
}

type DummyLocationService struct {
	locations []logistic.Location
	nextID    uint64
}

func NewDummyLocationService() *DummyLocationService {
	return &DummyLocationService{
		locations: initialLocations,
		nextID:    uint64(len(initialLocations) + 1),
	}
}

func (s *DummyLocationService) Describe(locationID uint64) (*logistic.Location, error) {
	for i := range s.locations {
		if s.locations[i].ID == locationID {
			return &s.locations[i], nil
		}
	}

	return nil, &LocationNotFoundError{locationID}
}

func (s *DummyLocationService) List(cursor uint64, limit uint64) ([]logistic.Location, error) {
	length := uint64(len(s.locations))

	if length == 0 {
		return nil, fmt.Errorf("nothing to list")
	}

	if cursor >= length {
		return nil, fmt.Errorf("cursor %d is out of range [0, %d]", cursor, length-1)
	}

	end := cursor + limit
	if end > length || limit == 0 {
		end = length
	}

	return s.locations[cursor:end], nil
}

func (s *DummyLocationService) Create(location logistic.Location) (uint64, error) {
	location.ID = s.nextID
	s.nextID++
	s.locations = append(s.locations, location)

	return location.ID, nil // TODO: Что делать с ошибкой?
}

func (s *DummyLocationService) Update(locationID uint64, location logistic.Location) error {
	for i := range s.locations {
		if s.locations[i].ID == locationID {
			s.locations[i] = location
			s.locations[i].ID = locationID
			return nil
		}
	}

	return &LocationNotFoundError{locationID}
}

func (s *DummyLocationService) Remove(locationID uint64) (bool, error) {
	for i := range s.locations {
		if s.locations[i].ID == locationID {
			s.locations = append(s.locations[:i], s.locations[i+1:]...)
			return true, nil // TODO: Зачем возвращать и bool, и error?
		}
	}

	return false, &LocationNotFoundError{locationID}
}

type LocationNotFoundError struct {
	ID uint64
}

func (e *LocationNotFoundError) Error() string {
	return fmt.Sprintf("location with ID %d not found", e.ID)
}

var initialLocations = []logistic.Location{
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
