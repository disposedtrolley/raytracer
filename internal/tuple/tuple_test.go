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
	tests := []struct{
		Name string
		FloatA float64
		FloatB float64
		ExpectedEqual bool
	}{
		{
			Name: "when two floats with a difference smaller than tuple.EPSILON are compared",
			FloatA: 0.000001,
			FloatB: 0.000002,
			ExpectedEqual: true,

		},
		{
			Name: "when two floats with a difference larger than tuple.EPSILON are compared",
			FloatA: 0.00001,
			FloatB: 0.0002,
			ExpectedEqual: false,
		},
		{
			Name: "when two floats with a difference equal to tuple.EPSILON are compared",
			FloatA: 0.00001,
			FloatB: 0.00002,
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
	tests := []struct{
		Name string
		TupleA *tuple.Tuple
		TupleB *tuple.Tuple
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