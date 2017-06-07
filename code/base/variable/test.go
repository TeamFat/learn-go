package main

import "fmt"

func main() {
	data:=[4]byte{65,66,3,4}
	for _,v := range data {
		fmt.Println(string(v))
	}
}