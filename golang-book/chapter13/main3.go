package main


import ("fmt";"flag";"math/rand")

func main() {
	// Define flags
	//fmt.Println(flag.Args())

	maxp := flag.Int("min", 0, "the max value")
	fmt.Println(*maxp)
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}