package main

import "fmt"
import "time"

func show() {

	for {
		fmt.Print("child ");
		time.Sleep(10000)
	}
}

func main() {

	go show()

	for {
		fmt.Print("parent ")
		time.Sleep(10000)
	}
}
