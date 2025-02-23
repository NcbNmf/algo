/*

01背包：每件物品最多只用一次
完全背包：每件物品有无限个
多重背包：每个物品的数量不一样
分组背包：N组物品，每组物品中有物品n个，每组中只能选一个物品

// 01背包

f[i][j] 表示只看前i个物品，总体积要求不超过j的情况下，总价值最大是多少
1、不选第i个物品：f[i][j] = f[i-1][j]
2、选第i个物品：f[i][j] = f[i-1][j-v[i]] + w[i]

然后不断递归f[i][j] = max(1, 2)就是答案了

1和2共同表示f[i][j]可以分成两部分
1表示当不选第i个物品时，从前i个物品中选择不超过体积j的物品，这时候因为没有比较所以还不是最大值
2表示当选择了第i个物品时，从前i个物品中选择不超过体积j-v[i]的物品，再加上第i个物品的w

之后在不断递归比较中就能得出f[i][j]的最大值
for i := 1; i <= n; i++ {
	for j := 1; j <= m; j++ {
		f[i][j] = f[i-1][j]
		if j >= v[i] {
			f[i][j] = max(f[i][j], f[i-1][j-v[i]] + w[i])
		}
	}
}

核心问题1：如何限制的体积
推一遍后可以发现
1、当i=1时，体积j迭代，但f[1][1~m]不变
2、当i=2时，体积j迭代，会有max(f[1][1~m], f[2][1~m])
这时候会有可以放下、放不下、只能放下1或2，三种情况

核心问题2：如何max只选择1或只选择2的f[][]呢
或许可以从后往前考虑
比如完全的：max(f[i-1][j], f[i-1][j-v[i]] + w[i])
当i=2时
一边是没有选择2的，也就是前面f[1][j]
一边是选择了2的，也就是后面的f[1][j-v[2]]+w[2]
这时候如果j-v[2]不能放下1(初始化的时候已经为0),那么max比对出来的就是只取了1和只取了2的

*/

package main
import (
	"os"
	"bufio"
	"fmt"
)

const N int = 1010

var (
	n	int
	m	int
	v	[N]int
	w	[N]int
	f	[N][N]int
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &v[i], &w[i])
	}

	//初始化
	for i := 0; i <= m; i++ {
		f[0][i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if v[i] <= j {
				f[i][j] = max(f[i][j], f[i-1][j-v[i]] + w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}