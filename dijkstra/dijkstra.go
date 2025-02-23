package main
import (
	"fmt"
	"os"
	"bufio"
)

const N int = 510
var n, m int		//n点m边
var g		[N][N]int	//邻接矩阵，如g[a][b]表示a指向b有一条边
var dist 	[N]int		//点n距离点1多远
var st		[N]bool		//记录这个点是否已经更新

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	//首先初始化距离
	//每个点一开始和1的距离都为无穷远
	//在后续的更新中，不断的比对来查询最短路
	for i := 1; i <= n; i++ {
		dist[i] = 0x3f3f3f3f	
	}
	dist[1] = 0

	//其次初始化邻接矩阵的每个边的长度
	//一开始也是指向无穷远
	//但是两点间存在链接的
	//后续会更新他们的边长
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			g[i][j] = 0x3f3f3f3f
		}
	}
	for i := 1; i <= m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		g[a][b] = min(g[a][b], c)	//这里即能够更新边长也能够去除重边和环
									//保留最短边
	}

	fmt.Println(dijkstra())
}

func dijkstra() int {
	for i := 0; i < n; i++ {
		t := -1	//每次循环都标记该点未更新
		for j := 1; j <= n; j++ {
			if st[j] == false && (t == -1 || dist[t] > dist[j]) {
				//t和st[]的作用类似
				//但是st[]更多是判断之前的点有没有更新过
				//t则是标记当前判断到的点是否为未更新的最小距离点
				//比如12345已经更新过了，那么前面的st[]就过不去
				//然后到了6，st[6]为更新，并且t==-1,那就进入这个点进行更新
				//后续再比对该点和其他与1有链接的点，谁的距离更小
				//更小的就更新先进入下一层判断
				t = j
			}
		}
		//承上启下
		//如果说，1与很多点有链接，但是只有少数几个点能够连接到后面的点n
		//并且这些连接不到点n的点边长都比较小，那t肯定会优先进入到最小的那个j阿
		//那改怎么办
		//没关系，因为上面的for会不断判断st[]，只要没更新过，那么下一层都会再进入再判断
		//一直到所有的点都更新过，并通过下面的dist[]=min(dist[], dist[]+g[][])来判断最小值

		for j := 1; j <= n; j++ {
			//这里的判断的意义为
			//上一层的t进入后，该t为与1距离最近的点
			//那么进来后会判断往后和1有链接的点，谁更小
			//因为dist[j]为1与j的距离，dist[t]为1与t的距离，g[t][j]为t与j的距离
			//为什么强调有链接，因为会加上一个g[][]，没有链接这个值就是0x3f3f3f3f
			//所以能够维持dist[j]的实质意义
			dist[j] = min(dist[j], dist[t]+g[t][j])
		}
		st[t] = true
	}
	
	if dist[n] == 0x3f3f3f3f {
		return -1
	}
	return dist[n]

	
}