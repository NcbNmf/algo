package main 
import (
	"fmt"
)

func main() {
	const N = 1010
	var n, m, q int
	var s [N][N]int
	var a [N][N]int

	fmt.Scanf("%d %d %d", &n, &m, &q)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + a[i][j]
		}
	}
	//这个for相当于把所有的s[i][j]都给算出来了
	//后续出现s[][]的时候就是进入到这里
	//而且因为输入的时候是i和j都等于1，所以i-1和j-1不会出现边界问题
	var x1, y1, x2, y2 int
	for q > 0 {
		fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
		fmt.Println(s[x2][y2] - s[x2][y1-1] - s[x1-1][y2] + s[x1-1][y1-1])
		q--
	}
	//s[x2][y2] - s[x2][y1-1] - s[x1-1][y2] + s[x1-1][y1-1]这就是算子矩阵的和
	return

}