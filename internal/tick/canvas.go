package tick

import (
	"fmt"
	"github.com/disposedtrolley/raytracer/internal/canvas"
	"github.com/disposedtrolley/raytracer/internal/colour"
	"github.com/disposedtrolley/raytracer/internal/tuple"
	"io/ioutil"
)

func RunCanvas() {
	velocityMultiplier := 11.25

	p := &Projectile{
		tuple.NewPoint(0, 1, 0),
		tuple.Mul(tuple.Normalise(tuple.NewVector(1, 1.8, 0)), velocityMultiplier),
	}

	e := &Environment{
		tuple.NewVector(0, -0.1, 0),
		tuple.NewVector(-0.01, 0, 0),
	}

	c := canvas.NewCanvas(900, 550)
	i := 0
	for p.Position.Y() > 0 {
		p = tick(e, p)
		c.Set(int(p.Position.X()), int(p.Position.Y()), colour.NewColour(255, 0, 0))
		fmt.Printf("Tick: %d Position: [%f, %f, %f]\n", i, p.Position.X(), p.Position.Y(), p.Position.Z())
		i++
	}

	out := canvas.ToPPM(c)

	_ = ioutil.WriteFile("./tick.ppm", []byte(out), 0644)
}