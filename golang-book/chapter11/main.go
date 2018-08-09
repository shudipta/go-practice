package main

import (
	"fmt"
	m "Go-Practice/golang-book/chapter11/Math"
)

func main() {
	xs := []float64{1,2,3,4}
	avg := m.Avg(xs)
	fmt.Println(avg)
	//defer func() {
		str := recover()
		fmt.Println("----", str)

	//}()
	panic("asdf")

}