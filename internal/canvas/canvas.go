package canvas

import (
	"fmt"
	"math"

	"github.com/disposedtrolley/raytracer/internal/colour"
	"github.com/disposedtrolley/raytracer/internal/tuple"
)

type Canvas struct {
	Data          []tuple.Tuple
	Width, Height int
}

func NewCanvas(width, height int) *Canvas {
	d := make([]tuple.Tuple, width*height)

	for i := 0; i < width*height; i++ {
		d[i] = colour.NewColour(0, 0, 0)
	}

	return &Canvas{
		Data:   d,
		Width:  width,
		Height: height,
	}
}

func (c *Canvas) Set(x, y int, pixel tuple.Tuple) {
	c.Data[c.Width*y+x] = pixel
}

func (c *Canvas) Get(x, y int) tuple.Tuple {
	return c.Data[c.Width*y+x]
}

const (
	ppmMaxLineLength  = 70
	ppmMaxColourValue = 255
)

func ToPPM(c *Canvas) (lines []string) {
	// header
	lines = []string{
		"P3",
		fmt.Sprintf("%d %d", c.Width, c.Height),
		fmt.Sprintf("%d", ppmMaxColourValue),
	}

	pixelsBuffered := 0
	buf := ""

	for _, pixel := range c.Data {
		rgb := colour.ToRGB(pixel)
		currTuple := fmt.Sprintf("%d %d %d", scale(rgb[0]), scale(rgb[1]), scale(rgb[2]))

		if pixelsBuffered+1 == c.Width || len(buf+currTuple) >= ppmMaxLineLength-1 {
			fmt.Printf("flushing at chars %d\n", len(buf+currTuple))
			// flush
			lines = append(lines, fmt.Sprintf("%s %s", buf, currTuple))

			buf = ""
			pixelsBuffered = 0
			continue
		}

		if len(buf) == 0 {
			buf = fmt.Sprintf("%s%s", buf, currTuple)
		} else {
			buf = fmt.Sprintf("%s %s", buf, currTuple)
		}

		pixelsBuffered++
	}

	lines = append(lines, "\n")
	return lines
}

func scale(val float64) int {
	if val < 0 {
		return 0
	} else if val > 1 {
		return 255
	} else {
		return int(math.Ceil(255 * (val / 1)))
	}
}
