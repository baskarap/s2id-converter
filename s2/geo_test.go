package s2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGeo_ReturnsNonNilGeoObject(t *testing.T) {
	geo := NewGeo(1)

	assert.NotNil(t, geo)
	assert.IsType(t, Geo{}, *geo)
}
