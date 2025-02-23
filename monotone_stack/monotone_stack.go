package main 

import (
	"fmt"
	"os"
	"bufio"
)

const N = 100010

func main() {

	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	var stack = make([]int, N)
	var tt int //栈顶指针

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)	

		for tt != 0 && stack[tt] >= x {
			tt--	//这里是判断，数x是否能进栈
				//如果原本的栈顶stack[tt]小于x，说明x可以进
				//跳出for，停止t--
		}
		if tt != 0 {
			fmt.Printf("%d ", stack[tt])	//这里就是跳出for后，直接输出当前栈顶
									//因为经过之前的判断，x肯定刚好大于当前栈顶
		
		} else {
			fmt.Printf("%d ", "-1")
		}
		tt++
		stack[tt] = x	//之前对比栈顶和x谁大，x都是在外面
						//这里就是把符合单调递增的x丢进栈里进行下一论判断
	}
}