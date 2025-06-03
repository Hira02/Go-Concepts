package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func (b Book) summary() {
	fmt.Printf("Book name: %s, writen by %s with %d pages\n", b.title, b.author, b.pages)
}
func (b *Book) _summary() {
	fmt.Printf("Book name: %s, writen by %s with %d pages\n", b.title, b.author, b.pages)
}
func normalFunc(a, b int) {
	fmt.Println("This is a normal function")
	fmt.Printf("%d + %d = %d", a, b, a+b)
}

func main() {
	b := Book{"Go Concept", "Hira", 200}
	fmt.Println(b)
	b.summary()
	bp := &b
	bp.author = "Hira Saha"
	b._summary()
	normalFunc(10, 20)

}
