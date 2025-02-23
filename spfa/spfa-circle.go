package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	index    int //记录节点编号
	distance int //记录该点到1的最短距离
}

var (
	n, m    int //节点数和边长
	a, b, c int

	h   []int //链表头
	e   []int
	ne  []int
	idx int
	w   []int
	//多加一个边权，边权的意义为：上一个点到这个点的距离
	//因为是稀疏图，所以不用g[][]，直接用邻接表

	inf = 0x3f3f3f3f

	st   []bool //记录这个点是否更新
	dist []int  //记录点n到1有多远

	heap []pair //堆中记录节点编号和到其路径长？
	size int    //堆大小

	cnt []int //求负环需要多加一个记录最短边数量的数组
	//因为如果存在环，xxxxxxxxxxxxx
)

// 邻接表构建
func add(a, b, c int) {
	e[idx] = b
	ne[idx] = h[a]
	w[idx] = c //a到b的距离，或者说上一个点到这个点的距离
	h[a] = idx
	idx++
}

func spfa() bool {
	for i := 1; i <= n; i++ {
		insert(pair{i, dist[i]})
		st[i] = true
	}
	//求负圈因为1可能到不了点n，所以需要把所有点都入堆
	//下面两行也就没有必要了
	// insert(pair{1, 0})
	// st[1] = true
	//这时候st存储的是st是否在堆内
	//后续会加一个判断，只有是更新过的点才会入堆
	for size > 0 {
		x := pop()
		st[x.index] = false //取出点后，标记为false，说明已不在堆内
		for i := h[x.index]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[x.index]+w[i] {
				dist[j] = dist[x.index] + w[i]
				//这里的判断是，只有当点j不在堆内且更新过j的最短距离时才入堆
				cnt[j] = cnt[x.index] + 1
				if cnt[j] >= n {
					return true
				}
				//为什么cnt[]+=1可以求负环
				//因为
				if st[j] == false {
					insert(pair{j, dist[j]})
					st[j] = true
				}
			}
		}
	}
	//不存在cnt[]>=n的情况
	return false
}

// /////////////////////
// heap operation
// 这里加入的堆排序，是针对边长的排序，是为了代替朴素做法中的第一个for
func insert(x pair) {
	size++
	heap[size] = x
	up(size)
}

// 删除操作，也是弹出
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
		tmp = x * 2
	}
	if x*2+1 <= size && heap[x*2+1].distance < heap[tmp].distance {
		tmp = x*2 + 1
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

// /////////////////////
// main
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
	cnt = make([]int, n+1)
	dist = make([]int, n+1)
	//求负环不需要dist无限大
	//只需要都指向一个虚拟源点，距离为0就可以，因为如果存在负权边
	//那在spfa()里也能判断变小
	// for i := 1; i <= n; i++ {
	//     dist[i] = inf
	// }
	// dist[1] = 0

	heap = make([]pair, max(n+1, m+1))

	//读入所有边
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b, &c)
		add(a, b, c)
	}

	if spfa() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
