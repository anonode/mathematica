package main

import (
	f "fmt"
	m "math"
)

var itercount int = 0

func sqrt(x float64) float64 {
	var (
		z_curr, z_last, epsilon float64 = x, x / 2, 0.00000001
	)
	for m.Abs(z_last-z_curr) > epsilon {
		itercount++
		z_last = z_curr
		z_curr = z_curr - (m.Pow(z_curr, 2)-x)/(2*z_curr) // could also be written as: z_curr - (z*z - x)/(2*z)
	}
	f.Println("Took", itercount, "iterations to find root.")
	return z_curr
}

func main() {
	var num float64
	f.Print("Find square root of: ")
	f.Scan(&num)
	f.Println(sqrt(num))
}
