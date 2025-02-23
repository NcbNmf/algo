package main
import (
	"fmt"
	"os"
	"bufio"
)

const N int = 100010
var (
	e	[N]int
	ne	[N]int
	h	[N]int
	idx	int
	n, m int
	q	[N]int
	d	[N]int	//记录入度
)

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}
func topusort() {
	hh, tt := 0, -1
	for i := 1; i <= n; i++ {
		if d[i] == 0 {
			tt++
			q[tt] = i	//i就是数x，x(i)入队
		}
	}

	//下面为什么不能用len(q) > 0 因为q是要输出的，不能做成动态
	//用指针tt，hh来表示会比较好
	for hh <= tt {	//队列不为空
		t := q[hh]	//取出队头
		hh++
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]	//找到出边(点)
			//下一个点的入度-1
			d[j]--
			if d[j] == 0 {
				tt++
				q[tt] = j	//入度都减没了，那就进入队列，放到下一轮判断
			}
		}
	}
	//如果队尾为n-1，说明一共进入了n-1个点
	//为什么是n-1，因为最后一个点进去就不会tt++了，因为进不到for循环
	if tt == n-1 {
		for i := 0; i <= tt; i++ {
			fmt.Printf("%d ", q[i])
		}
	} else {
		fmt.Println(-1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)

	//因为数据范围是N
	//所以初始化也要考虑N个链表头
	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		add(a, b)
		d[b]++	//a往b链接，b入度+1
	}

	topusort()

}

