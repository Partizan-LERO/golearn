package main

import "fmt"

func main() {
	defer theFour()
	defer theThree()
	defer theTwo()
	theOne()
}

func theOne()  {
	fmt.Println("The one")
}

func theTwo()  {
	fmt.Println("The two")
}

func theThree()  {
	fmt.Println("The three")
}

func theFour()  {
	fmt.Println("The four")
}