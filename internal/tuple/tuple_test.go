package tuple_test

import (
	"testing"

	"github.com/disposedtrolley/raytracer/internal/tuple"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	a := &tuple.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 1.0)

	assert.EqualValues(t,
		tuple.NewPoint(4, -4, 3),
		&tuple.Tuple{
			X: 4,
			Y: -4,
			Z: 3,
			W: 1.0,
		})
}

func TestVector(t *testing.T) {
	a := &tuple.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 0.0)

	assert.EqualValues(t,
		tuple.NewVector(4, -4, 3),
		&tuple.Tuple{
			X: 4,
			Y: -4,
			Z: 3,
			W: 0.0,
		})
}

func TestEqual(t *testing.T) {
	tests := []struct {
		Name          string
		FloatA        float64
		FloatB        float64
		ExpectedEqual bool
	}{
		{
			Name:          "when two floats with a difference smaller than tuple.EPSILON are compared",
			FloatA:        0.000001,
			FloatB:        0.000002,
			ExpectedEqual: true,
		},
		{
			Name:          "when two floats with a difference larger than tuple.EPSILON are compared",
			FloatA:        0.00001,
			FloatB:        0.0002,
			ExpectedEqual: false,
		},
		{
			Name:          "when two floats with a difference equal to tuple.EPSILON are compared",
			FloatA:        0.00001,
			FloatB:        0.00002,
			ExpectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tuple.Equal(tc.FloatA, tc.FloatB), tc.ExpectedEqual)
		})
	}

}

func TestEqualTuple(t *testing.T) {
	tests := []struct {
		Name          string
		TupleA        *tuple.Tuple
		TupleB        *tuple.Tuple
		ExpectedEqual bool
	}{
		{
			Name: "when two tuples with differences smaller than tuple.EPSILON are compared",
			TupleA: &tuple.Tuple{
				X: 0.000001,
				Y: 0.000001,
				Z: 0.000001,
				W: 0.0,
			},
			TupleB: &tuple.Tuple{
				X: 0.000002,
				Y: 0.000002,
				Z: 0.000002,
				W: 0.0,
			},
			ExpectedEqual: true,
		},
		{
			Name: "when two tuples with differences larger than tuple.EPSILON are compared",
			TupleA: &tuple.Tuple{
				X: 0.000001,
				Y: 0.000001,
				Z: 0.000001,
				W: 0.0,
			},
			TupleB: &tuple.Tuple{
				X: 0.00002,
				Y: 0.00002,
				Z: 0.00002,
				W: 0.0,
			},
			ExpectedEqual: false,
		},
		{
			Name: "when two tuples with differences equal to tuple.EPSILON are compared",
			TupleA: &tuple.Tuple{
				X: 0.00001,
				Y: 0.00001,
				Z: 0.00001,
				W: 0.0,
			},
			TupleB: &tuple.Tuple{
				X: 0.00002,
				Y: 0.00002,
				Z: 0.00002,
				W: 0.0,
			},
			ExpectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tuple.EqualTuple(tc.TupleA, tc.TupleB), tc.ExpectedEqual)
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		Name           string
		TupleA         *tuple.Tuple
		TupleB         *tuple.Tuple
		ExpectedResult *tuple.Tuple
	}{
		{
			Name:           "when two vectors are added",
			TupleA:         tuple.NewVector(5, -2, 3),
			TupleB:         tuple.NewVector(5, 2, -3),
			ExpectedResult: &tuple.Tuple{X: 10},
		},
		{
			Name:           "when a vector and a point are added",
			TupleA:         tuple.NewPoint(3, -2, 5),
			TupleB:         tuple.NewVector(-2, 3, 1),
			ExpectedResult: &tuple.Tuple{1, 1, 6, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedResult, tuple.Add(tc.TupleA, tc.TupleB))
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		Name           string
		TupleA         *tuple.Tuple
		TupleB         *tuple.Tuple
		ExpectedResult *tuple.Tuple
	}{
		{
			Name:           "subtracting two points should produce a vector",
			TupleA:         tuple.NewPoint(3, 2, 1),
			TupleB:         tuple.NewPoint(5, 6, 7),
			ExpectedResult: &tuple.Tuple{-2, -4, -6, 0},
		},
		{
			Name:           "subtracting a vector from a point should produce a point",
			TupleA:         tuple.NewPoint(3, 2, 1),
			TupleB:         tuple.NewVector(5, 6, 7),
			ExpectedResult: &tuple.Tuple{-2, -4, -6, 1},
		},
		{
			Name:           "subtracting two vectors should produce a vector",
			TupleA:         tuple.NewVector(3, 2, 1),
			TupleB:         tuple.NewVector(5, 6, 7),
			ExpectedResult: &tuple.Tuple{-2, -4, -6, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedResult, tuple.Sub(tc.TupleA, tc.TupleB))
		})
	}
}

func TestNeg(t *testing.T) {
	tests := []struct {
		Name           string
		Tuple          *tuple.Tuple
		ExpectedResult *tuple.Tuple
	}{
		{
			Name:           "when a vector is negated",
			Tuple:          tuple.NewVector(1, -2, 3),
			ExpectedResult: &tuple.Tuple{-1, 2, -3, 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedResult, tuple.Neg(tc.Tuple))
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		Name           string
		Tuple          *tuple.Tuple
		Scalar         float64
		ExpectedResult *tuple.Tuple
	}{
		{
			Name:           "multiplying a tuple by a scalar",
			Tuple:          &tuple.Tuple{1, -2, 3, -4},
			Scalar:         3.5,
			ExpectedResult: &tuple.Tuple{3.5, -7, 10.5, -14},
		},
		{
			Name:           "multiplying a tuple by a fraction scalar",
			Tuple:          &tuple.Tuple{1, -2, 3, -4},
			Scalar:         0.5,
			ExpectedResult: &tuple.Tuple{0.5, -1, 1.5, -2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedResult, tuple.Mul(tc.Tuple, tc.Scalar))
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		Name           string
		Tuple          *tuple.Tuple
		Scalar         float64
		ExpectedResult *tuple.Tuple
	}{
		{
			Name:           "dividing a tuple by a scalar",
			Tuple:          &tuple.Tuple{1, -2, 3, -4},
			Scalar:         2,
			ExpectedResult: &tuple.Tuple{0.5, -1, 1.5, -2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedResult, tuple.Div(tc.Tuple, tc.Scalar))
		})
	}
}
