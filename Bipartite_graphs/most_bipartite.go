package main
import (
	"bufio"
	"fmt"
	"os"
)

const N int = 510
const M int = 100010

var (
	n1	int
	n2	int
	m	int
	e	[M]int	//存点要大于N，因为可能e和ne存的右子图的点
				//可能会有多个指向
	ne	[M]int
	h	[N]int
	idx	int
	st	[N]bool		//右边的点的已有匹配记录
	match	[N]int	//右边的点对应的点
)

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func find(x int) bool {
	for i := h[x]; i != -1; i = ne[i] {
		j := e[i]
		if st[j] == false {
			st[j] = true
			//如果右边点没有匹配或者匹配的左边的那个点还能find到
			//那就迭代回去，让左边继续找下一个匹配
			if match[j] == 0 || find(match[j]) {
				match[j] = x 
				return true
			}
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n1, &n2, &m)
	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for m > 0 {
		m--
		var a, b int
		fmt.Fscan(in, &a, &b)
		add(a, b)	//只需要存储左指向右就可以
	}

	var res int
	for i := 1; i <= n1; i++ {
		//每次重置（清空）匹配
		//重置原因：
		//st只是在遍历左边点并存在迭代的时候
		//提供记录作用，画个图就推一遍就比较形象了
		for j := 0; j < N; j++ {
			st[j] = false
		}
		if find(i) {
			res++
		}
	}

	fmt.Println(res)
}
