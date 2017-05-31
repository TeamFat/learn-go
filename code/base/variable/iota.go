package variable

import "fmt"

func main() {
	//关键字iota定义常量数组中从0开始自增枚举值
	const(
		a = iota
		b
		c
		d
	)
	fmt.Println(a,b,c,d)

	//同一个常量组中可以提供多个iota，它们各自增长

	const(
		e,f = iota,iota
		g,h
		i,j
	)
	fmt.Println(g,h)

	//如果iota自增被打断，须显示恢复
	const(
		aa = iota
		bb = "bb"
		cc = "cc"
		dd = iota
	)
	fmt.Println(dd) //dd进行了显示恢复iota的值，计算的时候算上了bb,cc两行
}