package main

import "fmt"

//const MAX int = 10  // just for testing
const MAX int = 1000

func main() {
	var sum int = 0

	for i := 0; i < MAX; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println(i)
			sum = sum + i
		}
	}

	fmt.Printf("\n")
	fmt.Printf("anser is: %d\n", sum)
	fmt.Printf("done\n")
}
