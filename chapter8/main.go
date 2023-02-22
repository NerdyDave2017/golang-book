package main

import "fmt"

var y = 4

// Pointers
func main() {
	x := 5
	z := 20

	pointerFunc(&z)
	zero(x, y)
	fmt.Println(x, y, z)

	squarePointer(&x)

	fmt.Println(x)

}

func zero(x, y int) {
	x = 0
	y = 0
}

func pointerFunc(input *int) {
	*input = 0
}

func squarePointer(x *int) {
	*x = *x * *x
}
