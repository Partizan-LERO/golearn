package main

import "fmt"

func main() {
	var array[5] int
	array[0] = 100
	array[1] = 200
	array[2] = 300
	array[3] = 400
	array[4] = 500
	fmt.Println(array)

	for i:=0; i< len(array); i++ {
		fmt.Println(array[i])
	}

	for i, v := range array {
		fmt.Printf("key %v value %v", i, v)
		fmt.Println(" ")
	}

	var arrayTwo = [...] string {"a", "b"}
	fmt.Println(arrayTwo)

	var arrayThree[2] string
	arrayThree[0] = "c"
	arrayThree[1] = "d"

	fmt.Println(arrayThree == arrayTwo)
}
