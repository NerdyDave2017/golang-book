package main

import "fmt"

// Variables
var name = "Segun"

func main() {

	var class string

	class = "Golang"

	class = class + " is fun"

	animal := "Dog"

	animal = animal + " is a pet"

	fmt.Println(animal)

	fmt.Println(class)
	var myStr string = "Hello World"
	fmt.Println(myStr)
	var x string
	var y string = "Time"
	x = y + "table"
	fmt.Println(x)
	name := "Dayo"
	fmt.Println(name)
	myInt := 39
	fmt.Println(myInt)

	fmt.Print("Enter a Number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)

}

func testfunction() {
	fmt.Println(name)
}
