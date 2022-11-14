package main

import "fmt"

func main() {
	fmt.Println(half(34))
}

func half(x int) (half int, y bool) {
	half = x / 2
	y = x%2 == 0
	return
}
