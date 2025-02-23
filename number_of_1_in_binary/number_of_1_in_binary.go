package main

import (
	"fmt"
)

func lowbite(x int) int {
	return x & -x
}

func main() {

	var n int
	fmt.Scanf("%d", &n)
	var x int
	// var res int //res不能放在外面，会累加

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		var res = 0
		for x != 0 {
			x -= lowbite(x)
			res++
		}
		fmt.Print(res, " ")	//这里不能用println，因为是自动换行
	}

	
}