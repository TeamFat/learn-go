package main

import "fmt"

func main() {
	var i int
	for {
		fmt.Println(i)
		i++
		if i > 2 {

				goto TEST
		}
	}

	TEST:
	fmt.Println("break")
}