package canvas

import "github.com/disposedtrolley/raytracer/internal/tuple"

type Canvas struct {
	data []tuple.Tuple
	w, h int
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		data: make([]tuple.Tuple, width*height),
		w:    width,
		h:    height,
	}
}

func (c *Canvas) Set(x, y int, pixel tuple.Tuple) {
	c.data[c.w*y+x] = pixel
}

func (c *Canvas) Get(x, y int) tuple.Tuple {
	return c.data[c.w*y+x]
}
