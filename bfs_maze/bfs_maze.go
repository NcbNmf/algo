package main
import (
	"fmt"
	"bufio"
	"os"
)

type pair struct {
	first, second int
}

const N int = 110
var n, m int
var (
	g [N][N]int		//记录地图
	d [N][N]int		//记录距离
)

func bfs() int {
	var q = make([]pair, N)
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}		//定义xy轴的向量
	q = append(q, pair{0, 0})		//队头为第一个点
	
	for len(q) > 0 {
		t := q[0]					//取出队头
		q = q[1:]
		a, b := t.first, t.second
		for i := 0; i < 4; i++ {
			//4就是四个方向都判断一次
			//如果有两个方向可以走？
			//那就记录两个点到队列，一直到队列取不出为止
			//而且不会有两条路
			//因为先到达，也就是最短的路程会把路变成0，也就是走不出来
			lx := a + dx[i]
			ly := b + dy[i]
			if lx >= 0 && lx < n && ly >= 0 && ly < m && g[lx][ly] == 0 && d[lx][ly] == -1 {
				d[lx][ly] = d[a][b] + 1 	//记录距离+1 ，并且令经过的d[][]不为-1，也就是堵上
				q = append(q, pair{lx, ly})	//记录该点
			}
		}
	}
	return d[n-1][m-1]

}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &g[i][j])
			d[i][j] = -1
			//记录地图同时初始化距离记录
		}
	}
	d[0][0] = 0

	fmt.Println(bfs())

}