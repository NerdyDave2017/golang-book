package main

import "fmt"

//Arrays, Slices, Maps

func main() {
	// arrays()
	// slices()
	maps()
}

func arrays() {
	// var arr [5]int

	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// arr[3] = 4
	// arr[4] = 5

	arr := [5]int{1, 2, 3, 4, 5}

	var total float64 = 0
	// for i := 0; i < len(arr); i++ {
	// 	total += float64(arr[i])
	// }
	for _, value := range arr {
		total += float64(value)
	}
	fmt.Println(total)
}

func slices() {
	slices := []int{1, 2, 3, 5}
	x := append(slices, 4, 6)
	fmt.Println(x, slices)

	slice1 := []int{10, 25, 33, 27}
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	arr := []float64{23, 12, 25, 52, 84, 83}
	newArr := arr[0:]

	fmt.Println(newArr)
}

func maps() {
	// var map1 map[string]int
	// map1["age"] = 10

	// fmt.Println(map1) // This shows a run time error cos maps have to be initialized

	map1 := make(map[string]int)

	map1["age"] = 59

	fmt.Println(map1)
	fmt.Println(map1["age"])

	age, ok := map1["age"]

	fmt.Println(age, ok)

	if _, ok := map1["name"]; ok {
		fmt.Println("name field exist")
	} else {
		fmt.Println("name field doesn't exist")
	}

	// Shorter way to make maps

	map2 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
	}

	stateAndCapital := map[string]string{
		"Abia":    "Umahia",
		"Adamawa": "Yola",
	}

	fmt.Println(map2)
	fmt.Println(map2["H"])

	var x map[string]int
	fmt.Println(x)

	map3 := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"State": "Gas",
		},
	}

	fmt.Println(stateAndCapital)
	fmt.Println(map3)

}
