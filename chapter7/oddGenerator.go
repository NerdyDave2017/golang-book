package main

import "fmt"

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}

func makeOddGenerator() func() uint {

	i := uint(1)

	return func() (odd uint) {
		odd = i
		i += 2
		return
	}
}
