package main

import "fmt"

type EDocument struct {
	Number string
	Date string
	NumberOfPages int
	FooterDoc Footer
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
	doc2.FooterDoc.Author = "Anton"

	doc3 := new(EDocument)
	doc3.Date = "17-11-2018"
	doc3.Number = "0003A"
	doc3.NumberOfPages = 7
	doc3.FooterDoc.Author = "Иван"


	fmt.Printf("Type - %T, Value %v\n", doc1, doc1)
	fmt.Printf("Type - %T, Value %v\n", doc2, doc2)
	fmt.Printf("Type - %T, Value %v\n", doc3, doc3)

	doc1.showAuthor()
	doc2.showAuthor()
	doc3.showAuthor()

	doc1.showNumberOfPages()
	doc2.showNumberOfPages()
	doc3.showNumberOfPages()

	doc1.clearNumberOfPages()
	doc2.clearNumberOfPages()
	doc3.clearNumberOfPages()

	doc1.setPrefixAuthor()
	doc2.setPrefixAuthor()
	doc3.setPrefixAuthor()
}

func (doc EDocument) showAuthor(){
	fmt.Println(doc.FooterDoc.Author)
}

func (doc *EDocument) showNumberOfPages() {
	fmt.Println(doc.NumberOfPages)
}

func (doc *EDocument)  clearNumberOfPages() {
	doc.NumberOfPages = 0

	fmt.Println(doc.NumberOfPages)
}

func (doc *EDocument) setPrefixAuthor() {
	doc.FooterDoc.Author = "the Author is " + doc.FooterDoc.Author
	fmt.Println(doc.FooterDoc.Author)
}