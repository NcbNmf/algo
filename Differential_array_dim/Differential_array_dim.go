package main 
import (
	"fmt"
)

const N = 1010

var n, m, q int
var a, b [N][N]int

func insert(x1, y1, x2, y2 int, c int) {
	b[x1][y1] += c
	b[x2+1][y1] -= c
	b[x1][y2+1] -= c
	b[x2+1][y2+1] += c
}

func main() {

	fmt.Scanf("%d%d%d", &n, &m, &q)
	
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			insert(i, j, i, j, a[i][j])	//ijij,因为这里相当于
										//是一个空格一个空格的插入
		}
	}

	for q > 0 {
		var x1, y1, x2, y2, c int
		fmt.Scanf("%d%d%d%d%d", &x1, &y1, &x2, &y2, &c)
		insert(x1, y1, x2, y2, c)
		q--
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] += b[i-1][j] + b[i][j-1] - b[i-1][j-1]	//自身+=为前缀和，变成a
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Println()
	}
	return
}