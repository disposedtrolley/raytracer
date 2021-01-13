package tuple

import "math"

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

type Point Tuple

func NewPoint(x, y, z float64) *Point {
	return &Point{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}

type Vector Tuple

func NewVector(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
		W: 0.0,
	}
}

const EPSILON = 0.00001

// Equal returns whether a and b are equal to the tolerance
// defined in the constant EPSILON.
func Equal(a, b float64) bool {
	return math.Abs(a - b) < EPSILON
}