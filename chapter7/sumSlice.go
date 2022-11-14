package main

import "fmt"

// Sum of slice
func main() {
	sumOfSlice := func(input []int) (totalsum int) {
		for _, value := range input {
			totalsum += value
		}
		return
	}
	fmt.Println(sumOfSlice([]int{1, 2, 3, 4, 5}))
}
