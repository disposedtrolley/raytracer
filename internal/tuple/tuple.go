package tuple

import (
	"fmt"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}

func NewVector(x, y, z float64) *Tuple {
	return &Tuple{
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
	return math.Abs(a-b) < EPSILON
}

func EqualTuple(a, b *Tuple) bool {
	return Equal(a.X, b.X) &&
		Equal(a.Y, b.Y) &&
		Equal(a.Z, b.Z) &&
		Equal(a.W, b.W)
}

// Add returns a new Tuple which is the sum of a and b. Add will
// panic if a and b are both Point types, i.e. W = 1.
func Add(a, b *Tuple) *Tuple {
	if a.W+b.W > 1 {
		panic(fmt.Errorf("attempted to add two Points: %+v and %+v", a, b))
	}
	return &Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

// Sub returns a new Tuple which is the subtraction of b from a,
// i.e. a-b.
func Sub(a, b *Tuple) *Tuple {
	return &Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}
