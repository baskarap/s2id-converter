package s2

import (
	"errors"
)

const ErrInvalidLatlng = "Invalid LatLng being sent"

type Geo struct {
	level int
}

func (g *Geo) GetS2ID(lat float64, lng float64) (string, error) {
	return "", errors.New(ErrInvalidLatlng)
}

func NewGeo(level int) *Geo {
	return &Geo{
		level: level,
	}
}
