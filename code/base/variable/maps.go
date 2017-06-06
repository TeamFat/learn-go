package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//映射将键映射到值。
//映射的零值为 nil 。nil 映射既没有键，也不能添加键。
//make 函数会返回给定类型的映射，并将其初始化备用。

// var m map[string]Vertex

func main() {
	var m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//映射的文法与结构体相似，不过必须有键名。
	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)

	//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
	var o = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(o)
}
