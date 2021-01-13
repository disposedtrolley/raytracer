package tuple_test

import (
	"testing"

	"github.com/disposedtrolley/raytracer/internal/tuple"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	a := &tuple.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 1.0)

	assert.EqualValues(t,
		tuple.NewPoint(4, -4, 3),
		&tuple.Tuple{
			X: 4,
			Y: -4,
			Z: 3,
			W: 1.0,
		})
}

func TestVector(t *testing.T) {
	a := &tuple.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 0.0)

	assert.EqualValues(t,
		tuple.NewVector(4, -4, 3),
		&tuple.Tuple{
			X: 4,
			Y: -4,
			Z: 3,
			W: 0.0,
		})
}