package main

import "fmt"

func main() {
	sliceOfInt := []int{123,123,13,1,3,13,12,25,3,5}
	resSlice := calculate(sliceOfInt...)
	fmt.Println(resSlice)

	res := calculate(1,2,3,5,5,67,7,7,8)

	fmt.Println(res)
}

func calculate(numbers ...int) int  {
	result := 0

	for _, v := range numbers {
		result += v
	}

	return result
}