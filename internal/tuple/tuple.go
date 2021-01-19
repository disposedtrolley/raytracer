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

// Neg returns a new Tuple which is the negation of t.
func Neg(t *Tuple) *Tuple {
	// &Tuple{} is zero valued :)
	return Sub(&Tuple{}, t)
}

// Mul returns a new Tuple which is the product of t against
// the scalar value s.
func Mul(t *Tuple, s float64) *Tuple {
	return &Tuple{
		X: t.X * s,
		Y: t.Y * s,
		Z: t.Z * s,
		W: t.W * s,
	}
}

// Div returns a new Tuple which is the division of the scalar
// value s against t.
func Div(t *Tuple, s float64) *Tuple {
	return &Tuple{
		X: t.X / s,
		Y: t.Y / s,
		Z: t.Z / s,
		W: t.W / s,
	}
}

// Mag returns the magnitude of the vector tuple t.
func Mag(t *Tuple) float64 {
	if t.W != 0 {
		panic(fmt.Errorf("attempted to compute the magnitude of a non-vector tuple: %+v", t))
	}

	return math.Sqrt(
		math.Pow(t.X, 2) +
			math.Pow(t.Y, 2) +
			math.Pow(t.Z, 2) +
			math.Pow(t.W, 2))
}

func IsUnitVector(t *Tuple) bool {
	return Mag(t) == 1
}

func Normalise(t *Tuple) *Tuple {
	if t.W != 0 {
		panic(fmt.Errorf("attempted to normalise a non-vector tuple: %+v", t))
	}

	mag := Mag(t)

	return &Tuple{
		X: t.X / mag,
		Y: t.Y / mag,
		Z: t.Z / mag,
		W: t.W / mag,
	}
}
