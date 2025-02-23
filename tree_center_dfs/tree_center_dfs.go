package main
import (
	"fmt"
	"os"
	"bufio"
)
const N int = 100010
const M int = N * 2
var ans int = N

var (
	e	[M]int
	ne	[M]int
	h	[N]int	//n个链表的链表头
	idx	int
	n	int
	st	[N]bool	//记录哪些点已经遍历过
)

//构建一条a指向b的边
//相当于构建多次链表
//每有一个新的数被插，那被插的树就有新的链表头
func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(u int) int {
	//第一步，遍历以u为根节点的树有多少节点
	st[u] = true
	sum := 1
	res := 0	//
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]	//取得指向的点
		if st[j] == false {
			s := dfs(j)	//如果j没被搜过，那优先深度搜索
			
			res = max(res, s)
			//这里取极值比较好理解
			//s是u节点下面的某一子树的大小
			//sum是u节点加上子树的大小

			//这里为什么要比较呢？？？
			//res的作用更多是记录，s在不断+=之后会变得很大
			//或者说因为子树本来也是一个联通块，所以需要记录
			//这里就是记录子树连通块中点的数量
			
			sum += s
			//首先如果能够执行到这步
			//说明这是链表的倒数第二个点
			//返回的s肯定为1
			//所以sum这时候的+=是该点加上了下一个点
			//也就是该点加上他的子树一共有两个点
		}
	}
	// fmt.Println(res)
	//这一句可以直观的看到每次递推res的变化
	
	//这里为什么要比较呢？？？
	//和上面理由相似，这个是除去某节点后，剩下连通块中点的数量
	res = max(res, n-sum)
	//注意！！！！
	//这里的res不会返回上层循环
	//这里的res最主要的作用是不断和ans比对得出答案
	//sum能返回是因为有了一个return sum

	//因此在ans在不断记录中，res总会出现一次最小值，ans就记录这个最小值然后
	//输出
	ans = min(ans, res)
	return sum
}

func main() {
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < N; i++ {
		h[i] = -1	//初始化链表头
	}
	fmt.Fscan(in, &n)
	var a, b int
	for i := 0; i < n - 1; i++ {
	    fmt.Fscan(in, &a, &b)
		add(a, b)
		add(b, a)
	}
	dfs(1)
	fmt.Println(ans)
}