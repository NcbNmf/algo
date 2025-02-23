package main
import (
	"fmt"
	"bufio"
	"os"
)


const N int = 100010
var n, m int
var (
	d	[N]int		//记录距离
	h	[N]int		//n个链表链表头
	e	[N]int		//每个节点的值
	ne	[N]int		//每个节点的next
	idx	int
)

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func bfs() int {
	var q = make([]int, N)
	q = append(q, 1)		//队头为第一个点,在这题中就是1
	
	for len(q) > 0 {
		t := q[0]					//取出队头
		q = q[1:]
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if d[j] == -1 {
				d[j] = d[t] + 1 	//记录距离+1 ，并且令经过的d[]不为-1，也就是堵上
				q = append(q, e[i])	//记录该点
			}
		}
	}
	return d[n]

}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	for i := 0; i < N; i++ {
			h[i] = -1	//初始化表头
			d[i] = -1	//初始化遍历与否记录
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		add(a, b)
	}
	
	d[1] = 0	//这里表示是从1开始查找的

	fmt.Println(bfs())

}