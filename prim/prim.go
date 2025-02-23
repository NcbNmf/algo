package main
import (
	"fmt"
	"os"
	"bufio"
)

const inf int = 0x3f3f3f3f
var (
	n, m		int
	g			[][]int	//稠密图，用临界矩阵，就是点数^2≈边数
	dist		[]int	//距离存储的是被选中的点到集合的距离
	st			[]bool	//标记是否在集合内
	a, b, c		int
	res			int		//集合内所有边的和
)

func prim() int {
	for i := 0; i < n; i++ {
		t := -1	//t表示还未找到任何一个点
		for	j := 1; j <= n; j++ {
			if st[j] == false && (t == -1 || dist[t] > dist[j]) {
				//翻译：
				//如果j没在集合而且，当前没有找到任何一个或者
				//找到的这个点到集合的距离大于未在集合内的j到集合的距离
				//那就令t=j锁定j这个点
				//进行完迭代后，就能够找到一个集合外距离集合最近的点了
				t = j
			}
		}
		if i != 0 && dist[t] == inf {
			return inf
		}
		//翻译：
		//如果i不是第一个点并且迭代后找到的距离ihe最近的点的距离为inf
		//那说明图无法连通，也就没有最小生成树
		if i != 0 {
			res += dist[t]
		}
		st[t] = true	//然后把点计入集合

		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], g[t][j])
		}
		//翻译：
		//这一步是用点t更新集合外的点到集合的最小距离
		//比如一个点x，x→1为5，x→t为2，那么dist一开始存的dist[x]就是5
		//更新到t后，t和x有连接，并且距离还比x→1要短，那就更新dist[x]
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
    st = make([]bool, n+1)
	dist = make([]int, n+1)
	
	//二维切片make法
	for i := 0; i < n+1; i++ {
		tmp := make([]int, n+1)
		g = append(g, tmp)
	}
	//初始化距离
	for i := 0; i < n+1; i++ {
		dist[i] = inf
		for j := 0; j < n+1; j++ {
		    g[i][j] = inf
		}
	}
	

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b, &c)
		g[a][b], g[b][a] = min(g[a][b], c), min(g[a][b], c)	//无向图取最小边
	}

	res := prim()
	if res == inf {
		fmt.Println("impossible")
	} else {
		fmt.Println(res)
	}
}