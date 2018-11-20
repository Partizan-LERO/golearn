package main

import "fmt"

func main() {
	x := 100
	fmt.Println(x)
	someFunc()

	someFunctTwo()
}

func someFunc()  {
	x := 150
	fmt.Println(x)
}

func someFunctTwo()  {
	test := "Ok"
	point := 4

	fmt.Println(test)
	fmt.Println(point)
}