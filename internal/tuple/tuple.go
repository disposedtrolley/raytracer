package tuple

import (
	"fmt"
	"math"
)

type Tuple []float64

func (t Tuple) X() float64 {
	return t[0]
}

func (t Tuple) Y() float64 {
	return t[1]
}

func (t Tuple) Z() float64 {
	return t[2]
}

func (t Tuple) W() float64 {
	return t[3]
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

const EPSILON = 0.00001

// Equal returns whether a and b are equal to the tolerance
// defined in the constant EPSILON.
func Equal(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func EqualTuple(a, b Tuple) bool {
	return Equal(a.X(), b.X()) &&
		Equal(a.Y(), b.Y()) &&
		Equal(a.Z(), b.Z()) &&
		Equal(a.W(), b.W())
}

// Add returns a new Tuple which is the sum of a and b. Add will
// panic if a and b are both Point types, i.e. W = 1.
func Add(a, b Tuple) Tuple {
	if a.W()+b.W() > 1 {
		panic(fmt.Errorf("attempted to add two Points: %+v and %+v", a, b))
	}
	return Tuple{
		a.X() + b.X(),
		a.Y() + b.Y(),
		a.Z() + b.Z(),
		a.W() + b.W(),
	}
}

// Sub returns a new Tuple which is the subtraction of b from a,
// i.e. a-b.
func Sub(a, b Tuple) Tuple {
	return Tuple{
		a.X() - b.X(),
		a.Y() - b.Y(),
		a.Z() - b.Z(),
		a.W() - b.W(),
	}
}

// Neg returns a new Tuple which is the negation of t.
func Neg(t Tuple) Tuple {
	return Sub(Tuple{0, 0, 0, 0}, t)
}

// Mul returns a new Tuple which is the product of t against
// the scalar value s.
func Mul(t Tuple, s float64) Tuple {
	return Tuple{
		t.X() * s,
		t.Y() * s,
		t.Z() * s,
		t.W() * s,
	}
}

// Div returns a new Tuple which is the division of the scalar
// value s against t.
func Div(t Tuple, s float64) Tuple {
	return Tuple{
		t.X() / s,
		t.Y() / s,
		t.Z() / s,
		t.W() / s,
	}
}

// Mag returns the magnitude of the vector tuple t.
func Mag(t Tuple) float64 {
	if t.W() != 0 {
		panic(fmt.Errorf("attempted to compute the magnitude of a non-vector tuple: %+v", t))
	}

	return math.Sqrt(
		math.Pow(t.X(), 2) +
			math.Pow(t.Y(), 2) +
			math.Pow(t.Z(), 2) +
			math.Pow(t.W(), 2))
}

func IsUnitVector(t Tuple) bool {
	return Mag(t) == 1
}

func Normalise(t Tuple) Tuple {
	if t.W() != 0 {
		panic(fmt.Errorf("attempted to normalise a non-vector tuple: %+v", t))
	}

	mag := Mag(t)

	return Tuple{
		t.X() / mag,
		t.Y() / mag,
		t.Z() / mag,
		t.W() / mag,
	}
}

func Dot(a, b Tuple) float64 {
	if a.W() != 0 || b.W() != 0 {
		panic(fmt.Errorf("attempted to compute dot product of a point tuple: a: %+v b: %+v", a, b))
	}

	return a.X()*b.X() +
		a.Y()*b.Y() +
		a.Z()*b.Z()
}

func Cross(a, b Tuple) Tuple {
	if a.W() != 0 || b.W() != 0 {
		panic(fmt.Errorf("attempted to compute cross product of a point tuple: a: %+v b: %+v", a, b))
	}

	return NewVector(
		a.Y()*b.Z()-a.Z()*b.Y(),
		a.Z()*b.X()-a.X()*b.Z(),
		a.X()*b.Y()-a.Y()*b.X())
}
