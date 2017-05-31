package error

import (
	"fmt"
	"github.com/pkg/errors"
)

//go编程有一个惯用法，即把error类型的结果作为函数结果列表的最后一员
func divide(dividend int,divisor int) (result int,err error){
	if divisor==0 {
		err = errors.New("divisor by zero")
		return
	}
	result = dividend/divisor
	return
}

func main() {
	fmt.Println(divide(5,3))
	fmt.Println(divide(5,0))
}
