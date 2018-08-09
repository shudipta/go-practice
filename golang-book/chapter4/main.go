package main

import "fmt"
//import "os"

var x int = 5

func main()  {
	//x := "Hello World"

	//var x = "Hello World"

	fmt.Scan(&x)

	x += 5
	x++

	fmt.Print(x)

	f()

	//os.Exit(0)
}

func f() {
	fmt.Println(x)
}
