package main

import (
	"fmt"

	"./math"
)

// Creating packages
func main() {
	xs := []float64{1, 2, 3, 4}

	avg := math.Average(xs)
	fmt.Println(avg, ": Average")

	min := math.Min(xs)
	fmt.Println(min, ": Minimum")

	max := math.Max(xs)
	fmt.Println(max, ": Maximum")
}
