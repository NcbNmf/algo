package main
import "fmt"

func main() {

	var n, m string	//这里扫字符串%s，直接len就能得出多少位
	fmt.Scanf("%s", &n)
	fmt.Scanf("%s", &m)

	var N, M []int
	for i := len(n) - 1; i >= 0; i-- {
		N = append(N, int(n[i] - '0'))	//int(n[i] - '0')这是能够将字符char转成int的语法
	}
	for i := len(m) - 1; i >= 0; i-- {	//逆序加入
		M = append(M, int(m[i] - '0'))	//int(n[i] - '0')这是能够将字符char转成int的语法
	}
	
	c := add(N, M)
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
	
}

func add(N, M []int) []int {
	var c []int
	t := 0 //进位
	for i := 0; i < len(N) || i < len(M); i++ {
		if i < len(N) {
			t += N[i]
		}
		if i < len(M) {
			t += M[i]
		}
		c = append(c, t % 10)	//t如果大于10，取模后就得余数，小于10就等于原数
		t /= 10	//这里再取一次模的原因是，进一位，小于10就等于0
	}
	if t != 0 {
		c = append(c, 1)	//这里是否有必要？因为上述for已经把t给除掉了
							//有，假设两个切片长度相同，最后一位相加超过10
							//虽然前面已经有t/=10，但是没有办法进循环了
							//这一步就是为了完成进一的
	}
	return c
}