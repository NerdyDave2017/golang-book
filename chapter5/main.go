package main

import "fmt"

// Control statements
func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")
		}
	}
	for i := 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

testloop:
	for ban, con := 25, 0; con <= ban || con > ban; con++ {
		switch {
		case con < ban:
			fmt.Println(con, "is less than", ban)
		case con == ban:
			fmt.Println(con, "is equal to", ban)
		case con > ban:
			fmt.Println(con, "is greater than", ban)
			break testloop
		}
	}

	type circle struct {
	}

	type circle1 interface {
	}

}
