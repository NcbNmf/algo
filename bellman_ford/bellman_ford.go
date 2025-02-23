package main
import (
	"fmt"
	"os"
	"bufio"
)
const N int = 510
const M int = 10010
const inf int = 0x3f3f3f3f

type edge struct {
	a, b, w int		//w为a→b的长度
}
var n, m, k int

var (
	dist	[N]int
	edges	[M]edge
	last	[N]int		
	//last储存进入下一迭代的被更新过的数组
	//因为bellman不会记录更新过的点，因此在每次遍历边的时候总会遍历完每一条
	//那么假设只用一个数组dist，那在第一个点A更新后，没迭代进下一层
	//那么基于这个点A的下一个点B只会更新A-B的距离，因为点1到点A的距离没被记录
	//理应最短值是2的，现在只有1
)
func bellman_ford() {
	for i := 0; i < k; i++ {
		copy(last[0:], dist[0:])	//没确定一次边就复制一次修改还过的数组
		for i := 0; i < m; i++ {
			e := edges[i]
			dist[e.b] = min(dist[e.b], last[e.a]+e.w)
		}
	}
	if dist[n] > inf/2 {	
		//为什么要/2呢？？？
		//因为可能存在负边，所以有可能dist小于inf，但小于inf/2不太可能
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n])
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &k)

	for i := 0; i < m; i++ {
		a, b, w := 0, 0, 0
		fmt.Fscan(reader, &a, &b, &w)
		edges[i] = edge{a, b, w}
	}

	for i := 1; i < N; i++ {
		dist[i] = inf
	}
	dist[1] = 0

	bellman_ford()
}