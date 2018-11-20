package main

import "fmt"

func main() {
	a, b, c := 10, 50, 100

	if a > b {
		fmt.Println("a")
	} else if b > c {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}
	
	var text = "string"

	switch text {
	case "string" :
		fmt.Println(text)
	fallthrough
	case "admin":
		fmt.Println("admin")
	default:
		fmt.Println("Error")
	}

	switch  {
	case text == "string" || text == "admin":
		fmt.Println(a)
	default:
		fmt.Println("default")
	}

	for x:=1; x < 100; x++ {
		if x % 10 == 0 {
			fmt.Println("x = ", x)
			continue
		}

		fmt.Println(x)
	}
}
