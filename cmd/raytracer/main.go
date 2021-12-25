package main

import (
	"fmt"

	"github.com/disposedtrolley/raytracer/internal/tick"
)

func main() {
	fmt.Println("Hello, raytracer")

	tick.Run()
	tick.RunCanvas()
}
