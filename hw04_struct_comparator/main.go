package main

import (
	"fmt"
)

type book struct {
	Id            int
	Title, Author string
	Year, Size    int
	Rate          float32
}

func (p *book) Print() {
	fmt.Println("Id:", p.Id)
	fmt.Println("Title:", p.Title)
	fmt.Println("Author:", p.Author)
	fmt.Println("Year:", p.Year)
	fmt.Println("Size:", p.Size)
	fmt.Println("Rate:", p.Rate)
}

func main() {
	book1 := &book{1, "skazki", "Pushkin", 1825, 23, 4.3}
	book2 := &book{2, "drama", "Lermontov", 1817, 302, 4.6}
	book1.Print()
	fmt.Println("--------------")
	book2.Print()
	fmt.Println("--------------")
	var s string

	fmt.Printf("Введите поле для сравнения Year, Size или Rate: ")
	fmt.Scanf("%s", &s)
	switch {
	case s == "Year" || s == "year":
		fmt.Println(book1.Year > book2.Year)
	case s == "Size" || s == "size":
		fmt.Println(book1.Size > book2.Size)
	case s == "Rate" || s == "rate":
		fmt.Println(book1.Rate > book2.Rate)
	}
}
