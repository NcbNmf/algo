package main
import (
	"fmt"
	"bufio"
	"os"
)

type ull int64	//这里的意思是，超过int64范围的会自动取模
				//因为哈希值是取模后才能得到嘛，这个2^63就是取模的那个值

var h []ull
var p []ull

const N = 100010
const P = 131 //自定进制，131或者13331

func getHash(l, r int) ull {
	return h[r] - h[l-1] * p[r-l+1]
}

func main() {
	h = make([]ull, N)
	p = make([]ull, N)

	in := bufio.NewReader(os.Stdin)
	var n, m int
	var str string
	fmt.Fscan(in, &n, &m, &str)
	str = " " + str		//调整下标从1开始

	//预处理p[]和h[]
	//p[]处理成，有n位数就有p^n
	//h[]处理成，前n前缀的哈希值，超过int64自动取模
	p[0] = 1
	for i := 1; i <= n; i++ {
		p[i] = p[i-1]*P
		h[i] = h[i-1]*P + ull(str[i])
	}

	for i := 0; i < m; i++ {
		var l1, r1, l2, r2 int 
		fmt.Fscan(in, &l1, &r1, &l2, &r2)
		if getHash(l1, r1) == getHash(l2, r2) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}