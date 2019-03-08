package s2

import (
	"errors"
	"github.com/golang/geo/s2"
)

const (
	ErrInvalidLatlng = "Invalid LatLng being sent"
	ErrInvalidLevel  = "Invalid level of S2Geo being sent"
)

type Geo struct {
	level int
}

func (g *Geo) GetS2ID(lat float64, lng float64) (string, error) {
	latLng := s2.LatLngFromDegrees(lat, lng)
	if !latLng.IsValid() {
		return "", errors.New(ErrInvalidLatlng)
	}

	return "", errors.New(ErrInvalidLevel)
}

func NewGeo(level int) *Geo {
	return &Geo{
		level: level,
	}
}
