package main

import "fmt"

func main() {
	someWord := scanSomeWords()
	fmt.Println(someWord)
	fmt.Println("someWord")
}

func scanSomeWords() string {
	var someWords string
	fmt.Scan(&someWords)
	return someWords
}
