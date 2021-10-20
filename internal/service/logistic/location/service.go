package location

import (
	"errors"
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
		locations: logistic.InitialLocations,
		nextID:    uint64(len(logistic.InitialLocations) + 1),
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
		return nil, ErrEmptyList
	}

	if cursor >= length {
		return nil, fmt.Errorf("cursor %d is out of range [0, %d]", cursor, length-1)
	}

	var err error = nil
	end := cursor + limit
	if end >= length || limit == 0 {
		end = length
		err = ErrEndOfList
	}

	return s.locations[cursor:end], err
}

func (s *DummyLocationService) Create(location logistic.Location) (uint64, error) {
	if err := location.Validate(); err != nil {
		return 0, err
	}

	location.ID = s.nextID
	s.nextID++
	s.locations = append(s.locations, location)

	return location.ID, nil
}

func (s *DummyLocationService) Update(locationID uint64, location logistic.Location) error {
	if err := location.Validate(); err != nil {
		return err
	}

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

var (
	ErrEmptyList = errors.New("empty locations list")
	ErrEndOfList = errors.New("no more locations in list")
)

type LocationNotFoundError struct {
	ID uint64
}

func (e *LocationNotFoundError) Error() string {
	return fmt.Sprintf("location with ID %d not found", e.ID)
}
