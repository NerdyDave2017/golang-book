package main

import "fmt"

//  Find the smallest number in the list

func main() {
	x := []int64{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := getSmallest(x)
	fmt.Println(smallest)

	fmt.Println(x)
	fmt.Println(sortArr(x))
}

func getSmallest(x []int64) int64 {
	pivot := x[0]
	for _, value := range x {
		switch {
		case value < pivot:
			pivot = value
		}
	}

	return pivot
}

// @todo finish sorting algorithm
func sortArr(x []int64) []int64 {

	for i, j := 0, 1; i < len(x)-1; i, j = j, j+1 {
		if x[i] > x[j] {
			x[i], x[j] = x[j], x[i]
		}
	}
	return x
}
