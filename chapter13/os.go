package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open file Long method
	readFile()
	// Open file Short method
	readFile2()
	// Create file
	createFile()
}

func readFile() {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// get the file size
	stat, err := file.Stat()

	if err != nil {
		fmt.Println((err))
		return
	}

	// read the file
	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)

	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func readFile2() {
	bs, err := ioutil.ReadFile("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs)

	fmt.Println(str)
}

func createFile() {
	file, err := os.Create("testing.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString("This is a testing file")
}
