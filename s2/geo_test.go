package s2_test

import (
	"github.com/baskarap/s2id-converter/s2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGeo_ReturnsNonNilGeoObject(t *testing.T) {
	geo := s2.NewGeo(1)

	assert.NotNil(t, geo)
	assert.IsType(t, s2.Geo{}, *geo)
}

func TestGetS2ID_ReturnsError_WhenInvalidLatlongIsGiven(t *testing.T) {
	invalidLatLng := 190.0
	geo := s2.NewGeo(1)
	s2id, err := geo.GetS2ID(invalidLatLng, invalidLatLng)

	assert.Empty(t, s2id)
	assert.NotNil(t, err)
}
