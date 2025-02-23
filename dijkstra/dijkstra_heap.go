package main
import (
	"os"
	"fmt"
	"bufio"
)

type pair struct {
	index int		//记录节点编号
	distance int	//记录该点到1的最短距离
}

var (
	n, m	int		//节点数和边长
	a, b, c int

	h		[]int	//链表头
	e		[]int
	ne		[]int
	idx		int
	w		[]int	
	//多加一个边权，边权的意义为：上一个点到这个点的距离
	//因为是稀疏图，所以不用g[][]，直接用邻接表

	inf = 0x3f3f3f3f

	st		[]bool	//记录这个点是否更新
	dist	[]int	//记录点n到1有多远

	heap	[]pair	//堆中记录节点编号和到其路径长？
	size	int		//堆大小
)

//邻接表构建
func add(a, b, c int) {
	e[idx] = b
	ne[idx] = h[a]
	w[idx] = c	//a到b的距离，或者说上一个点到这个点的距离
	h[a] = idx
	idx++
}

func dijkstra() int {
	insert(pair{1, 0})
	for size > 0 {
		x := pop()
		if x.index == n {
			return x.distance
		}
		//s[]是标记某点是否已更新
		//和朴素dijkstra中的t:=-1是一样的
		if st[x.index] {
			continue
		}
		//////////////
		/*
		上面的代码很好的说明了为什么堆模拟的dijkstra
		比朴素的好
		因为在朴素中，上面的应该是不断比对n个s[]是否已更新
		同时还要判断dist[]的最小点，在for中要操作n次
		但这个dist[]的最小点在堆的up和down中只需要操作log(n)
		*/
		//////////////
		st[x.index] = true

		//随后遍历邻接表
		for i := h[x.index]; i != -1; i = ne[i] {
			j := e[i]
			dist[j] = min(dist[j], dist[x.index] + w[i])
			insert(pair{j, dist[j]})
		}
		//////////////
		/*
		这里遍历邻接表也有优化
		原本是要遍历完n个点，但是邻接表的特性使得
		这一段只需要遍历有链接的点就可以

		h[x.index]为表头指向的下一个点的idx
		x.index代表表头
		e[i]为表头指向的下一个点的值
		w[i]代表上一个点到这个点的距离
		*/
		//////////////
	}
	//到达不了n
	return -1
}

///////////////////////
//heap operation
//这里加入的堆排序，是针对边长的排序，是为了代替朴素做法中的第一个for
func insert(x pair) {
	size++
	heap[size] = x
	up(size)
}
//删除操作，也是弹出
func pop() pair {
	x := heap[1]
	heap[1] = heap[size]	
	size--
	down(1)
	return x
}
func down(x int) {
	tmp := x
	if x*2 <= size && heap[x*2].distance < heap[tmp].distance {
		tmp = x*2
	}
	if x*2+1 <= size && heap[x*2+1].distance < heap[tmp].distance {
		tmp = x*2+1
	}
	if x != tmp {
		heap[x], heap[tmp] = heap[tmp], heap[x]
		down(tmp)
	}
}
func up(x int) {
	for x/2 > 0 && heap[x/2].distance > heap[x].distance {
		heap[x], heap[x/2] = heap[x/2], heap[x]
		x /= 2
	}
}
///////////////////////
//main
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	//初始化链表头
	h = make([]int, n+1)
	for i := 1; i <= n; i++ {
        h[i] = -1
    }
    e, ne, w = make([]int, m), make([]int, m), make([]int, m)
    
	//初始化距离
	st = make([]bool, n+1)
    dist = make([]int, n+1)
	for i := 1; i <= n; i++ {
        dist[i] = inf
    }
	dist[1] = 0
	
	heap = make([]pair, max(n+1, m+1))
	
	//读入所有边
	for i := 0; i < m; i++ {
        fmt.Fscan(reader, &a, &b, &c)
        add(a, b, c)
    }

	fmt.Println(dijkstra())

}

