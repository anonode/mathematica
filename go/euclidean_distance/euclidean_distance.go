package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func makevertex(ex, why float64) vertex {
	v := vertex{}
	v.x = ex
	v.y = why
	return v // could have also done this: `return vertex{ex,why}`
}

func euclid_dist(c1, c2 vertex) float64 {
	return math.Sqrt(math.Pow(c2.y-c1.y, 2) + math.Pow(c2.x-c1.x, 2))
}

func main() {
	var x, y float64
	fmt.Println("Enter first coordinate: ")
	fmt.Scanln(&x, &y)
	v1 := makevertex(x, y)
	fmt.Println("Enter second coordinate: ")
	fmt.Scanln(&x, &y)
	v2 := makevertex(x, y)
	dist := euclid_dist(v1, v2)
	fmt.Println("Euclidean distance is: ", euclid_dist(v1, v2))
	if dist == math.Trunc(dist) {
		fmt.Println("\nYour coordinates formed a pythagorean triple!\n" +
			"This means that the sides of the triangle formed" +
			"by your coordinates are integers!")
	}
}
