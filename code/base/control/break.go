package main

import "fmt"

//直接跳出所有嵌套循环， 这时候就可以用到go的label breaks特征了
func main() {
	L1:
		for i := 0;i<3;i++{
			L2:
				for j := 0;j<5 ;j++  {
					if(j>2){
						continue L2
					}
					if(i>1){
						break L1
					}
					fmt.Println(i, ":", j, " ");
					
				}
		}
}
