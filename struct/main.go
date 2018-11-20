package main

import "fmt"

type EDocument struct {
	Number string
	Date string
	NumberOfPages int
	Footer
}

type Footer struct {
	Author string
}

func main() {
	doc1 := EDocument{
		"0001A",
		"19-11-2018",
		10,
		Footer{"Sergey"},
	}

	var doc2 EDocument

	doc2.Date  = "18-11-2018"
	doc2.Number = "0002A"
	doc2.NumberOfPages = 11
	doc2.Author = "Anton"

	fmt.Printf("Type - %T, Value %v\n", doc1, doc1)
	fmt.Printf("Type - %T, Value %v", doc2, doc2)
}
