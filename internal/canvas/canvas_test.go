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
