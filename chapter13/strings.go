package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaaa", "a", "b", 2),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
	)

	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})

	fmt.Println(arr, str)

	str1 := "ab£"
	chars := []rune(str1)
	for i := 0; i < len(chars); i++ {
		
		char := string(chars[i])
		println(char)
	}

}
