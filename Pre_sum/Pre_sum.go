package main 
import (
	"fmt"
)

func main() {

	var n, m int
	var a, b int

	fmt.Scanf("%d %d", &n, &m)
	A := make([]int, n+1)
	S := make([]int, n+1)	//切片记得make，记得n+1
	A[0] = 0

	for i := 1; i <= n; i++ {	//这里要<=n，要不就会少扫一个
		fmt.Scanf("%d", &A[i])
	}

	for i := 1; i <= n; i++ {
		S[i] = S[i - 1] + A[i]	//相当于递归出一个函数
	}

	for m > 0 {
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(S[b] - S[a - 1])
		m--
	}

}