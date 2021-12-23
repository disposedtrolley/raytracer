package colour_test

import (
	"testing"

	"github.com/disposedtrolley/raytracer/internal/colour"
	"github.com/disposedtrolley/raytracer/internal/tuple"
	"github.com/stretchr/testify/assert"
)

func TestColourOperations(t *testing.T) {
	t.Run("adding colours", func(t *testing.T) {
		c1 := colour.NewColour(0.9, 0.6, 0.75)
		c2 := colour.NewColour(0.7, 0.1, 0.25)

		assert.Equal(t, colour.NewColour(1.6, 0.7, 1.0),
			tuple.Add(c1, c2))
	})

	t.Run("subtracting colours", func(t *testing.T) {
		c1 := colour.NewColour(0.9, 0.6, 0.75)
		c2 := colour.NewColour(0.7, 0.1, 0.25)

		assert.True(t, tuple.EqualTuple(
			tuple.Sub(c1, c2),
			colour.NewColour(0.2, 0.5, 0.5)))
	})

	t.Run("multiplying a colour by a scalar", func(t *testing.T) {
		c1 := colour.NewColour(1, 0.2, 0.4)
		c2 := colour.NewColour(0.9, 1, 0.1)

		assert.True(t, tuple.EqualTuple(
			colour.HadamardProduct(c1, c2),
			colour.NewColour(0.9, 0.2, 0.04)))
	})
}
