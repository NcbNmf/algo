package main
import "fmt"

func main() {

	var a string
	var b int	//这里是把b看作一个整体去乘

	fmt.Scanf("%s", &a)
	fmt.Scanf("%d", &b)
	if b == 0 {
		fmt.Println(0)
		return
	}

	var A []int
	for i := len(a)-1; i >= 0; i-- {
		A = append(A, int(a[i] - '0'))
	}
	C := mul(A, b)
	for i := len(C)-1; i >= 0; i-- {
		fmt.Printf("%d", C[i])
	}
}

func mul(A []int, b int) []int {
	var t int
	var C []int
	for i := 0; i < len(A); i++ {
		t += A[i] * b 
		C = append(C, t % 10)
		t /= 10
	}
	//下面的for很好的解决了，t多出来的情况
	//无论是多出1位数还是2位数的t，都能用这个for来解决
	for t > 0 {
		C = append(C, t % 10)
		t /= 10
	}
	return C
}