package main
import (
	"fmt"
	"os"
	"bufio"
)

const N int = 211
const inf int = 0x3f3f3f3f
var (
	d	[N][N]int	//矩阵d[i][j]记录i-j的距离
	n, m, Q	int
	a, b, c	int
)

func floyd() {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
} 

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &Q)

	//初始化距离
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				d[i][j] = 0 	//这里判断是去除自环
			} else {
			    d[i][j] = inf   //注意不能放在外面，要不上面if执行后，下面又执行，会被覆盖
			}
			
		}
	}

	//加入边
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b, &c)
		d[a][b] = min(d[a][b], c)	//这里求最小值的目的是去除重边
	}

	floyd()

	for Q > 0 {
		Q--
		var x, y int
		fmt.Fscan(reader, &x, &y)
		if d[x][y] >= inf/2 {
			fmt.Println("impossible")
		} else {
			fmt.Println(d[x][y])
		}
	}
}
