package control

import "fmt"


//switch灵活多变，还可以用于类型判断
//if和switch都可以包含一条初始化语句
func main() {
	x := []int{1, 2, 3}
	i := 3
	switch i {
	case x[1]:
		fmt.Println("a")
	case x[2], 1:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}


	a := 10
	switch a {
	case 10:
		println("aa")
		fallthrough //需要继续下一分支，使用fallthrough
	case 0:
		println("cc")
	case 120:
		println("dd")
	}

	b := []int{1, 2, 3}
	switch {
	case b[1] > 0:      //当做if else使用
		println("a")
	case b[1] < 0:
		println("b")
	default:
		println("c")
	}

	d := []int{1, 2, 3}
	switch i := d[2]; { //带初始值
	case i > 0:
		println("a")
	case i < 0:
		println("b")
	default:
		println("c")
	}

}