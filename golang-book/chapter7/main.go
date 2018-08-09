package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3rd")
}
func main() {
	defer third()
	defer second()
	first()
}
