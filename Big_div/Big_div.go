package main 
import (
	"fmt"
)

func main() {
	var a string
	var b int
	fmt.Scanf("%s", &a)
	fmt.Scanf("%d", &b)
	var A []int
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i] - '0'))
	}

	r := 0	//余数在外面定义使其可以输出
	c := div(A, b, &r)

	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
	fmt.Println()
	fmt.Println(r)
}

func div(A []int, b int, r *int) []int {
	var c []int
	*r = 0	//这里为什么要让r指向的值变成0呢，未知
	for i := len(A) - 1; i >= 0; i-- {
		*r = *r*10 + A[i]
		c = append(c, *r/b)	//这里直接用余数除以b，可以省略掉一些代码
		*r %= b
	}
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
        c[i], c[j] = c[j], c[i]
    }
	//为什么上面要把第一位换成最后一位
	//因为前缀零只能通过切片切切片的方式去掉
	//而c[:len(c)-1]这一句只有后面是开区间
	//因此要先把一个正序的c变成倒叙去掉0,再倒叙输出
	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]	//去掉前缀0
	}
	return c
}