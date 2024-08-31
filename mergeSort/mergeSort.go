package main

import (
	"fmt"
)

func main() {

	var (
		n int
	)
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	tmp := make([]int, n)

	merge_Sort(q, tmp, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}

func merge_Sort(q []int, tmp []int, left, right int) {

	if left >= right {
		return
	}
	mid := (left+right)/2
	merge_Sort(q, tmp, left, mid)
	merge_Sort(q, tmp, mid+1, right)

	//下面代码其实可以从最底层开始理解，也就是当对半划分到只剩下两个数之后
	//下面的代码一开始跑，就是有序的了
	k, l, r := 0, left, mid+1	//l为左半边起点，r为右半边起点

	for l <= mid && r <= right {
		if q[l] <= q[r] {
			tmp[k] = q[l]
			k++
			l++
		} else {
			tmp[k] = q[r]
			k++
			r++
		}
	}
	//下面判断左右两边是否已经遍历完
	for l <= mid {
		tmp[k] = q[l]
		k++
		l++
	}
	for r <= right {
		tmp[k] = q[r]
		k++
		r++
	}
	//上面两端的精妙之处在于，如果一段已经遍历完，那l或r必定大于mid或right
	//那么就不回执行for
	for i, j := left, 0; i <= right; i, j = i+1, j+1 {
        q[i] = tmp[j] // 把临时有序序列存回q序列中
    }
	//必须在这里执行，因为这里的第一次执行还是处于只有两个数的时候
	//如果不是在这里执行，那最后输出的时候就会很奇怪
	//而且必须是left，这是考虑到右边输入的left不同，不一定是0
}