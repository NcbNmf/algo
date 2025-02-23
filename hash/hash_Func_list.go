package main
import (
	"fmt"
	"os"
	"bufio"
)

const N int = 100003	//为什么这里不是100010，因为据大量实践和研究
						//用离数据范围最近的质数转换成的哈希值，冲突概率最小

var (
	h	[N]int
	e	[N]int
	ne	[N]int
	idx	int
)

func insert(x int) {
	k := (x % N + N ) % N	//这是为了把哈希值k维持在正数
							//例如假设x=-10,N=3，计算一下就知道
	//然后就是链表的头插法了
	e[idx] = x
	ne[idx] = h[k]
	h[k] = idx
	idx++
}

func find(x int) bool {
	k := (x % N + N ) % N	//先找到链表的头部
	for i := h[k]; i != -1; i = ne[i] {
		if e[i] == x {
			return true
		}
	}
	return false

}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < N; i++ {
		h[i] = -1
	}
	//h[]是用来定义表头的，因此它储存的值初始化为-1
	//而h[]的下标才是转化成的hash值

	for i := 0; i < n; i++ {
		var op string
		fmt.Fscan(in, &op)
		switch op {
		case "I":
			var x int
			fmt.Fscan(in, &x)
			insert(x)
		case "Q":
			var x int
			fmt.Fscan(in, &x)
			if find(x) {
			    fmt.Println("Yes")
			} else {
			    fmt.Println("No")
			}
		}
	}
}