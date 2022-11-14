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

}

func zero(x, y int) {
	x = 0
	y = 0
}

func pointerFunc(input *int) {
	*input = 0
}
