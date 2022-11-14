package main

import "fmt"

func main() {
	fmt.Println(largestNum(2, 55, 6, 3, 53, 635, 553, 3, 353535, 22, 335))

}

func largestNum(args ...int) (greatest int) {
	for _, value := range args {
		switch {
		case value > greatest:
			greatest = value
		}
	}
	return
}
