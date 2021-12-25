package canvas

import (
	"fmt"
	"math"
	"strconv"
	"strings"

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

func ToPPM(c *Canvas) (out string) {
	var b strings.Builder

	// header
	b.WriteString("P3\n")
	b.WriteString(fmt.Sprintf("%d %d\n", c.Width, c.Height))
	b.WriteString(fmt.Sprintf("%d\n", ppmMaxColourValue))

	channelsBuffered := 0
	buf := ""

	var pixelChannels []string
	for _, pixel := range c.Data {
		rgb := colour.ToRGB(pixel)
		for _, channel := range rgb {
			pixelChannels = append(pixelChannels, strconv.Itoa(scale(channel)))
		}
	}

	for _, channel := range pixelChannels {
		if len(buf) == 0 {
			buf = fmt.Sprintf("%s%s", buf, channel)
		} else {
			buf = fmt.Sprintf("%s %s", buf, channel)
		}

		channelsBuffered++

		if len(buf) + 3 >= ppmMaxLineLength {
			b.WriteString(buf)
			b.WriteString("\n")
			buf = ""
		} else if channelsBuffered/3 >= c.Width {
			b.WriteString(buf)
			b.WriteString("\n")
			buf = ""
			channelsBuffered = 0
		}
	}

	b.WriteString("\n")
	return b.String()
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
