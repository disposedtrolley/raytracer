package colour

import "github.com/disposedtrolley/raytracer/internal/tuple"

func NewColour(r, g, b float64) tuple.Tuple {
	return tuple.Tuple{r, g, b, 0}
}

func HadamardProduct(c1, c2 tuple.Tuple) tuple.Tuple {
	return NewColour(
		c1[0]*c2[0],
		c1[1]*c2[1],
		c1[2]*c2[2])
}

func ToRGB(c tuple.Tuple) []float64 {
	return []float64{c[0], c[1], c[2]}
}
