package main
import "fmt"

const N int = 10
var n int
var path	[N]int	//用来记录输出数列
var st		[N]bool	//用来记录数字能否使用

func dfs(x int) {
	if x == n {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println()
	}
	

	for i := 1; i <= n; i++ {
		if st[i] == false {
			path[x] = i
			st[i] = true
			dfs(x+1)
			st[i] = false
		}
	}
}

func main() {

	fmt.Scanf("%d", &n)
	
	dfs(0)
}