package main

import (
	"fmt"
)

type Book struct {
	id            int
	title, author string
	year, size    int
	rate          float32
}

type Num int

const (
	Year Num = iota
	Size
	Rate
)

func (p *Book) Print() {
	fmt.Println("Id:", p.id)
	fmt.Println("Title:", p.title)
	fmt.Println("Author:", p.author)
	fmt.Println("Year:", p.year)
	fmt.Println("Size:", p.size)
	fmt.Println("Rate:", p.rate)
}

func main() {
	book1 := &Book{1, "skazki", "Pushkin", 1825, 23, 4.3}
	book2 := &Book{2, "drama", "Lermontov", 1817, 302, 4.6}
	book1.Print()
	fmt.Println("--------------")
	book2.Print()

	var s string

	fmt.Printf("Введите поле для сравнения Year, Size или Rate: ")
	fmt.Scanf("%s", &s)
	switch {
	case s == "Year" || s == "year":
		fmt.Println(book1.year > book2.year)
	case s == "Size" || s == "size":
		fmt.Println(book1.size > book2.size)
	case s == "Rate" || s == "rate":
		fmt.Println(book1.rate > book2.rate)
	}
}
