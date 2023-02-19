package main

import "fmt"

// Data types

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[6])
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
