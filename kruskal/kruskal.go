package main
import (
	"fmt"
	"os"
	"bufio"
	"sort"
)

const N int = 100010
const inf = 0x3f3f3f3f

type Edge struct {
	a, b, w int
}

var (
	reader = bufio.NewReader(os.Stdin)
	n, m int
	e []Edge	//储存边
	p []int		//作为并查集的节点
				//如p[A]=A 意思是原本每个点的父节点都是自己
)

func find(x int) int {
	if x != p[x] {
		p[x] = find(p[x])
	}
	return p[x]
}

func kru() int {
	//sort 使用方法
	//排序的意义是为了最小生成树
	sort.Slice(e, func(i, j int) bool {
		return e[i].w < e[j].w
	})

	//为每一个节点赋予一个父节点，也就是它自己
	//后面结合find()就是找到祖宗节点
	for i := 1; i <= n; i++ {
		p[i] = i
	}

	var count, res int //count储存有多少条边加入，res储存边权和

	//用find找到祖宗节点
	//枚举每一条边，注意是边不是点
	//这里也注意一下a, b的变化，当find()过后
	//a，b就不是原本指向的点了
	for i := 0; i < m; i++ {
		a, b, w := e[i].a, e[i].b, e[i].w
		a, b = find(a), find(b)
		if a != b {
			p[a] = b
			res += w
			count++
		}
	}
	if count < n-1 {
		return inf 
	}
	return res

}

func main() {
	fmt.Fscan(reader, &n, &m)
    
    //注意这里e不能多开空间
    //因为后面sort是用e来排序的，多开空间就会多几个0
	e = make([]Edge, m)
	p = make([]int, n+10)

	//存入边和点
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &e[i].a, &e[i].b, &e[i].w)
	}

	res := kru()
	if res == inf {
		fmt.Println("impossible")
	} else {
		fmt.Println(res)
	}
}

