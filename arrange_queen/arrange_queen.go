package main
import "fmt"

const N int = 20
var n int
var path	[N][N]string	//用来记录皇后位置
var (
	col		[N]bool	//记录当前列
	dg		[N]bool //记录对角线1
	udg		[N]bool	//记录对角线2
)

func dfs(x int) {
	if x == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf(path[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
	}
	

	for i := 0; i < n; i++ {
		if !col[i] && !dg[i+x] && !udg[i-x+n]  {
			//上面的判断等价于false
			//也就是当变量为true时，说明被使用，同时加入！var，表示为false
			//false时if不通过
			//相当于被使用时不执行if

			path[x][i] = "Q"
			col[i], dg[i+x], udg[i-x+n] = true, true, true
			dfs(x+1)
			//后面两步就是恢复现场
			//除了true改false，填入的Q也变回.
			col[i], dg[i+x], udg[i-x+n] = false, false, false
			path[x][i] = "."
		}
	}
}

func main() {

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = "."
		}
	}
	
	dfs(0)

}