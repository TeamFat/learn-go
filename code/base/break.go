package main

import "fmt"

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
