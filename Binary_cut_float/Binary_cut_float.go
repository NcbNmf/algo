package main
import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scanf("%f", &x)

	var (
		l float64 = -10000
		r float64 = 10000
	)

	for r - l > 0.0000001 {
		//这里没有什么+1 -1 ，因为浮点数不需要考虑边界问题
		//只需要考虑保留到小数后多少位就行
		//也仍然还是采用逼近，二分的思想，判断mid是在x的左边还是右边
		mid := (l + r) / 2
		if mid * mid * mid >= x {
			r = mid
		} else {
			l = mid
		}
	}

	fmt.Printf("%f", l)
}