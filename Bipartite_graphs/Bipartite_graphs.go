package main
import (
	"fmt"
	"os"
	"bufio"
)

const N int = 100010
const M int = N * 2

var (
	e	[M]int
	ne	[M]int
	h	[N]int
	idx	int
	n	int
	m	int
	color	[N]int
	reader = bufio.NewReader(os.Stdin)
)

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func dfs(u int, c int) bool {
	//首先令第一个点染上颜色，是1 是2无所谓
	//因为dfs，确定了祖宗节点，那他的后续所有节点也都会确定颜色
	color[u] = c
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		if color[j] != 1 && color[j] != 2 {
			//上面语句首先判断j是否染过色
			//没有的话直接dfs深度遍历
			//直接去到最底部，染色
			//如果dfs过程中出现，染上同一个颜色
			//那就直接返回false
			if dfs(j, 3 - c) == false {
				return false
			}
		} else if color[j] == c {
		    return false
		}
	}
	//上面的连通块遍历完，没有问题那就返回true
	return true
}

func main() {
	fmt.Fscan(reader, &n, &m)
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	for i := 0; i < m; i++ {
	    var a, b int
	    fmt.Fscan(reader, &a, &b)
	    add(a, b)
	    add(b, a)
	}

	var flag = true

	for i := 1; i <= n; i++ {
		if color[i] != 1 && color[i] != 2 {
			//定义dfs返回bool
			//如果染色过程发生矛盾，返回false
			if dfs(i, 1) == false {
				flag = false
				break
			}
		}
	}

	if flag == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

