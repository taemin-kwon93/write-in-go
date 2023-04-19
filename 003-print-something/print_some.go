package main

import "fmt"

func main() {
	// strAndQuotedStr()
	scanSomeWord()
}

func strAndQuotedStr() {
	str := "Hello, World\nIt Good To See U!"
	fmt.Printf("%s", str)
	fmt.Printf("%q", str) // quoted string literal
}

func scanSomeWord() {
	var word string
	fmt.Scan(&word)
	fmt.Println(word)
}
