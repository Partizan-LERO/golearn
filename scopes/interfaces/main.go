package main

import "fmt"

type Document struct {
	Date string
	Number string
	NumberOfPages int
}

type PersonCard struct {
	Date string
	FirstName string
	LastName string
	Age int
}

type PrintInterface interface {
	CheckData()
}

func (d *Document) CheckData() {
	if d.Date != "" {
		fmt.Println("Date in doc is correct")
	} else {
		fmt.Println("Date in doc is incorrect")
	}
}

func (c *PersonCard) CheckData() {
	if c.Date != "" {
		fmt.Println("Date in card is correct")
	} else {
		fmt.Println("Date in card is incorrect")
	}
}

func main() {
	doc := new(Document)
	doc.Date = "12-12-2014"
	doc.Number = "00001A"
	doc.NumberOfPages = 12

	card := new(PersonCard)
	card.Date = "12-13-2014"
	card.FirstName = "Sergey"
	card.LastName = "Samoylenko"
	card.Age = 27

	docs := []PrintInterface{doc, card}
	printAnyDoc(docs)
}

func printAnyDoc(anyDoc []PrintInterface)  {
	for _, v := range anyDoc {
		fmt.Println(v)
	}
}