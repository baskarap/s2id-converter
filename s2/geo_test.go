package s2_test

import (
	"github.com/baskarap/s2id-converter/s2"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ValidLat = -6.244378
	ValidLng = 106.798477
)

func TestNewGeo_ReturnsNonNilGeoObject(t *testing.T) {
	geo := s2.NewGeo(1)

	assert.NotNil(t, geo)
	assert.IsType(t, s2.Geo{}, *geo)
}

func TestGetS2ID_ReturnsError_WhenInvalidLatLngIsGiven(t *testing.T) {
	invalidLatLng := 190.0
	geo := s2.NewGeo(1)
	s2id, err := geo.GetS2ID(invalidLatLng, invalidLatLng)

	assert.Empty(t, s2id)
	assert.Equal(t, s2.ErrInvalidLatlng, err.Error())
}

func TestGetS2ID_ReturnsError_WhenInvalidLevelIsBeingSet(t *testing.T) {
	geo := s2.NewGeo(99)
	s2id, err := geo.GetS2ID(ValidLat, ValidLng)

	assert.Empty(t, s2id)
	assert.Equal(t, s2.ErrInvalidLevel, err.Error())
}

func TestGetS2ID_ReturnsS2ID_WhenValidLatLngAndLevelAreGiven(t *testing.T) {
	geo := s2.NewGeo(13)
	s2id, err := geo.GetS2ID(ValidLat, ValidLng)

	assert.Equal(t, "3344469644458065920", s2id)
	assert.Nil(t, err)
}
