package canvas_test

import (
	"testing"

	"github.com/disposedtrolley/raytracer/internal/canvas"
	"github.com/disposedtrolley/raytracer/internal/colour"
	"github.com/stretchr/testify/assert"
)

func TestCanvasWrite(t *testing.T) {
	c := canvas.NewCanvas(10, 20)
	red := colour.NewColour(1, 0, 0)

	c.Set(2, 3, red)

	assert.Equal(t, c.Get(2, 3), red)
}

func TestToPPM(t *testing.T) {
	t.Run("converts a 5x3 canvas to a string in PPM format", func(t *testing.T) {
		c := canvas.NewCanvas(5, 3)
		c.Set(0, 0, colour.NewColour(1.5, 0, 0))
		c.Set(2, 1, colour.NewColour(0, 0.5, 0))
		c.Set(4, 2, colour.NewColour(-0.5, 0, 1))

		expected := []string{
			"P3",
			"5 3",
			"255",
			"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
			"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0",
			"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255",
			"\n",
		}

		assert.Equal(t, expected, canvas.ToPPM(c))
	})

	t.Run("splits long lines", func(t *testing.T) {
		width := 10
		height := 2
		c := canvas.NewCanvas(10, 2)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				c.Set(x, y, colour.NewColour(1, 0.8, 0.6))
			}
		}

		expected := []string{
			"P3",
			"10 2",
			"255",
			"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
			"153 255 204 153 255 204 153 255 204 153 255 204 153",
			"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
			"153 255 204 153 255 204 153 255 204 153 255 204 153",
		}

		assert.Equal(t, expected, canvas.ToPPM(c))
	})
}
