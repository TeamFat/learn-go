package main

//不过使用分组导入语句是更好的形式。
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	//在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。 例如， Pizza 就是个已导出名， Pi 也同样，它导出自 math 包。
	fmt.Println(math.Pi)
}
