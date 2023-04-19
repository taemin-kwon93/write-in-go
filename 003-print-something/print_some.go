package main

import "fmt"

func main() {
	strAndQuotedStr()
}

func strAndQuotedStr() {
	str := "Hello, World\nIt Good To See U!"
	fmt.Printf("%s", str)
	fmt.Printf("%q", str) // quoted string literal
}
