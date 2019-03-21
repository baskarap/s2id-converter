package s2

import (
	"errors"
	"fmt"
	"github.com/baskarap/s2id-converter"
	"github.com/golang/geo/s2"
)

const (
	ErrInvalidLatlng = "Invalid LatLng being sent"
	ErrInvalidLevel  = "Invalid level of S2Geo being sent"
)

var _ s2id_converter.Geo = (*Geo)(nil)

type Geo struct {
	level int
}

func (g *Geo) GetS2ID(lat float64, lng float64) (string, error) {
	latLng := s2.LatLngFromDegrees(lat, lng)
	if !latLng.IsValid() {
		return "", errors.New(ErrInvalidLatlng)
	}

	if !isValidLevel(g.level) {
		return "", errors.New(ErrInvalidLevel)
	}

	cell := s2.CellIDFromLatLng(latLng).Parent(g.level)
	return fmt.Sprintf("%d", cell), nil
}

func isValidLevel(level int) bool {
	if level <= 30 && level > -1 {
		return true
	}
	return false
}

func NewGeo(level int) *Geo {
	return &Geo{
		level: level,
	}
}
