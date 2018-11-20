package main

import "fmt"

func main() {
	a, b := 14, 35

	fmt.Println(sum(a, b))
	fmt.Println(question("question"))

	logic, number, str := multiple(a, b, 123)

	fmt.Println(logic)
	fmt.Println(number)
	fmt.Println(str)

	_, numberTwo, _ := multiple(a, b, 123)
	fmt.Println(numberTwo)
}

func sum(a int, b int) int  {
	return a + b
}

func question(s string) bool  {
	if s == "question" {
		return true
	} else {
		return false
	}
}


func multiple(a, b, c int) (bool, int, string) {
	if a > b {
		return true, a, "больше"
	} else {
		return false, c, "меньше"
	}
}