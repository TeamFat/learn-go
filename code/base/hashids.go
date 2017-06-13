package main

import "fmt"
import "github.com/speps/go-hashids"

func main() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}
