package main

import (
	"fmt"
)

func main() {
	var simpleMap = map[int]string{}

	simpleMap[198] = "error"

	fmt.Println(simpleMap)

	var anotherMap = make(map[string]int)

	anotherMap["one"] = 123

	fmt.Println(anotherMap)


	anotherMap["two"] = 1213
	anotherMap["three"] = 1213
	anotherMap["four"] = 1213

	fmt.Println(anotherMap)

	delete(anotherMap, "two")

	fmt.Println(anotherMap)

	if value, ok := anotherMap["one"]; ok {
		fmt.Println(value)
	}
}
