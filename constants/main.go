package main

import "fmt"

const myConst =  "const"

const (
	_ = iota
	_ = iota
	_ = iota
	D = iota + 5
	E = iota
	_ = iota
)

func main() {
	const constTwo = 123

	fmt.Println(myConst)
	fmt.Println(constTwo)

	var numberTwo int = 10

	result := constTwo + numberTwo

	fmt.Println(result)

	fmt.Println(D)
	fmt.Println(E)

	x, _, z := TypesGo()

	fmt.Println(x)
	fmt.Println(z)
}

func TypesGo() (int, bool, string) {
	return 100, false, "stringType"
}