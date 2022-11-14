package main

import "fmt"

// Fibonacci Number
func main() {
	fmt.Println(fib(7))
	fmt.Println(fibSequence(65))
}

func fib(x int) int {
	if x <= 1 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// Fibonacci Sequence
func fibSequence(x int) (fibSlice []int) {
	for i := 0; i <= x; i++ {
		nextFib := getFibNum(i)
		fibSlice = append(fibSlice, nextFib)
	}
	return
}

func getFibNum(x int) int {
	a := 0
	b := 1
	for i := 0; i < x; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}
