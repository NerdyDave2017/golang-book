package main

import "fmt"

// Convert feet to meter

func main() {
	fmt.Print("Enter value in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	meter := input * 0.3048

	fmt.Println(meter)
}
