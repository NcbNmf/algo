package main
import (
	"fmt"
	"os"
	"bufio"
)

const N int = 200003	//开放寻址法一般开2~3倍空间
						//同时也是找范围最近的质数

const null int = 0x3f3f3f3f //这个数的意义为绝对在值域外的数
							//用来表示这个下标没有储存值

var (
	h	[N]int
)

func find(x int) int {
	k := (x % N + N ) % N
	for h[k] != null && h[k] != x {	//首先是判断这个坑位有没有储存值
									//第二句有存在必要嘛？
									//因为for ()是无条件循环
									//所以加入的判断是，这个坑位不为空，而且坑位上不是同一个人
									//那就k++
									//如果为空那就直接退出返回k继续行了
		k++
		if k == N {
			k = 0	//这里特判是，如果在h[]找到头了还是没有找到坑位
					//回头从0开始找
		}
	}
	return k

}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < N; i++ {
		h[i] = null
	}
	//h[]是用来定义表头的，因此它储存的值初始化为-1
	//而h[]的下标才是转化成的hash值

	for i := 0; i < n; i++ {
		var op string
		fmt.Fscan(in, &op)
		switch op {
		case "I":
			var x , k int
			fmt.Fscan(in, &x)
			k = find(x)
			h[k] = x
		case "Q":
			var x int
			fmt.Fscan(in, &x)
			if h[find(x)] == x {
			    fmt.Println("Yes")
			} else {
			    fmt.Println("No")
			}
		}
	}
}