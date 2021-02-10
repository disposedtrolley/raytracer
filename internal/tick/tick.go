// Package tick implements the example at the end of Chapter 1.
package tick

import (
	"fmt"

	"github.com/disposedtrolley/raytracer/internal/tuple"
)

type Projectile struct {
	Position tuple.Tuple
	Velocity tuple.Tuple
}

type Environment struct {
	Gravity tuple.Tuple
	Wind    tuple.Tuple
}

func tick(environment *Environment, projectile *Projectile) *Projectile {
	return &Projectile{
		Position: tuple.Add(projectile.Position, projectile.Velocity),
		// TODO tuple.Add() and possibly other operations can have variadic arguments.
		Velocity: tuple.Add(tuple.Add(projectile.Velocity, environment.Gravity), environment.Wind),
	}
}

func Run() {
	velocityMultiplier := 2.0

	p := &Projectile{
		tuple.NewPoint(0, 1, 0),
		tuple.Mul(tuple.Normalise(tuple.NewVector(1, 1, 0)), velocityMultiplier),
	}

	e := &Environment{
		tuple.NewVector(0, -0.1, 0),
		tuple.NewVector(-0.01, 0, 0),
	}

	i := 0
	for p.Position.Y() > 0 {
		p = tick(e, p)
		fmt.Printf("Tick: %d Position: [%f, %f, %f]\n", i, p.Position.X(), p.Position.Y(), p.Position.Z())
		i++
	}
}
