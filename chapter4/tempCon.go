package main

import "fmt"

// Convert Fahrenheit to Celcius
func main() {
	fmt.Print("Enter Fahrenheit value: ")
	var input float64
	fmt.Scanf("%f", &input)
	c := (input - 32) * 5 / 9
	fmt.Println(c, "ºC")
}
