package main
import (
	"fmt"
)

const N = 100010
var a = make([]int, N)
var b = make([]int, N)

func insert(l, r int, c int) {
	b[l] += c
	b[r+1] -= c
}

func main() {
	
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 1; i <= n; i++ {
		insert(i, i, a[i])
	}
	//第二个for是为了构建b差分数组
	//how
	//推一下就知道了b[1]=a[1]，b[2]=a[2]-a[1]，b[3]=a[3]-a[2]
	//这样a[i]就等于b[1]+···+b[i]

	for m > 0 {
		var l, r, c int
		fmt.Scanf("%d%d%d", &l, &r, &c)
		insert(l, r, c)
		m--
		//这里就是insert的第二个作用，加c
	}

	for i := 1; i <= n; i++ {
		b[i] += b[i-1]
	}
	//这个for的作用是把b加了c后转成a，因为b就是a的前缀和

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", b[i])
	}

}