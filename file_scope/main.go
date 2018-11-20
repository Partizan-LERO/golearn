package main

import "fmt"

var car = "foo"
func main() {
	fmt.Println(car)
	checkSpareWheel()
}

func checkSpareWheel()  {
	if car == "foo" {
		fmt.Println("Spare wheel exists")
	}
}
