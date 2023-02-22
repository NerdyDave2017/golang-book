package main

import "fmt"

// Functions

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(sumOfTwo(2, 4))
	fmt.Println(checkArr(arr))
	fmt.Println(arguements(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

	// Closure, anonymous functions
	multiply := func(x int, y int) (result int) {
		result = x * y
		return
	}
	fmt.Println(multiply(2, 4))

	fmt.Println(factorial(5))
	deferTest()
	panicRecover()
}

func sumOfTwo(x int, y int) int {
	return x + y
}

func checkArr(x []int) ([]int, int) {
	return x, len(x)
}

func arguements(args ...int) []int {
	return args
}

func sum(args ...int) int {
	total := 0

	for _, value := range args {
		total += value
	}
	return total
}

// Recursion
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Defer
func deferTest() {
	defer func() {
		fmt.Println("Function1")
	}()

	func() {
		fmt.Println("Function2")
	}()
}

// Panic and recover

func panicRecover() {
	defer func() {
		str := recover()
		fmt.Println(str, "This is a recover message")
	}()
	panic("Error Message")

}
