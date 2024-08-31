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
	
	if cmp(N, M) {
		c := sub(N, M)
		for i := len(c) - 1; i >= 0; i-- {
			fmt.Printf("%d", c[i])
		}
	} else {
		fmt.Print("-")
		c := sub(M, N)
		for i := len(c) - 1; i >= 0; i-- {
			fmt.Printf("%d", c[i])
		}
	}
}

func sub(N, M []int) []int {
	var c []int
	t := 0 //进位
	for i := 0; i < len(N); i++ {
		t = N[i] + t	//先减一个t，t只有0 和 1 ，判断前面有没有借
						//这一句也是为了之后让t对M[i]操作
		if i < len(M) {
			t -= M[i]	//t = t - M[i]	如果没有借位那就是N[i]-M[i]，也会出t是否大于0
		}
		
		if t < 0 {
			c = append(c, (t + 10) % 10)
			t = -1
		} else {
			c = append(c, t % 10)
			t = 0
		}
		// c = append(c, (t + 10) % 10)	//其实这一句不用加判断，虽然后面t还是要加判断
										//因为如果t>=0，那等于t本身，如果t<0，那加10
	}
	//如果N第一位等于1，那可能会出现前导0
	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]	//切片切切片，从第0位切到倒数第二位，后面的那个]是)，开区间
	}
	
	return c
}

func cmp(A, B []int) bool {
	if len(A) != len(B) {
		return len(A) > len(B)	//这里其实是已经判断了，就是返回lenA是否大于lenB
	}
	//下面就是长度相等后，每个位上的大小判断
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] != B[i] {
			return A[i] > B[i]
		}
	}
	return true //这是相等时候的情况
}