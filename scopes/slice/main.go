package main

import "fmt"

func main() {
	var sliceOne []int
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	sliceOne = append(sliceOne,10)
	sliceOne = append(sliceOne,20)
	sliceOne = append(sliceOne,30)
	sliceOne = append(sliceOne,40)
	sliceOne = append(sliceOne,50)

	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))


	fmt.Println(sliceOne[len(sliceOne) - 2:])
	fmt.Println(sliceOne[3:])
	fmt.Println(sliceOne[:3])
	fmt.Println(sliceOne[2:4])


	slice :=  make([]int, 5, 19)
	fmt.Println(slice)
}
