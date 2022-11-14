package main

import "fmt"

//  Find the smallest number in the list

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := getSmallest(x)
	fmt.Println(smallest)
}

func getSmallest(x []int) int {
	pivot := x[0]
	for _, value := range x {
		switch {
		case value < pivot:
			pivot = value
		}
	}

	return pivot
}
