package main

import "fmt"

func main() {
	sliceOne := []int{123,12315,54,3,5,5}
	res := calculate(14)

	a := res(123131)
	fmt.Println(a)


	x := someFunc(sliceOne, func(i int) int {
		return i + 1000
	})

	fmt.Println(x)
}


func calculate(i int) func(int) int {
	fmt.Println(i)
	
	return func(k int) int {
		return k * 1000
	}
}

func sum(a int) int  {
	return a * 1000
}


func someFunc(numbers []int, justFunc func(int) int) int {
	result := 0
	for _, v := range numbers {
		fmt.Println(justFunc(v))
		result += justFunc(v)
	}

	return result
}