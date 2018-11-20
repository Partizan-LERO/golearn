package main

import "fmt"

func main() {
	fmt.Println(someFunc(6, 4))

	fmt.Println(closure("My text"))
}

func someFunc(a, b int) int {
	z := a + b

	func(text string) {
		fmt.Println(text)
	}("Some text")

	return z
}


func closure(s string) string {
	text := "sentence " + s

	result := func() string {
		return text + " some text"
	}

	return result()
}