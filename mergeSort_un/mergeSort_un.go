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

	fmt.Println(merge_Sort(q, tmp, 0, n-1))
}

func merge_Sort(q []int, tmp []int, left, right int) int {

	if left >= right {
		return 0
	}
	mid := (left+right)/2

	res := merge_Sort(q, tmp, left, mid) + merge_Sort(q, tmp, mid+1, right)


	k, l, r := 0, left, mid+1
	
	for l <= mid && r <= right {
		if q[l] <= q[r] {
			tmp[k] = q[l]
			k++
			l++
		} else {
			tmp[k] = q[r]
			k++
			r++
			res += mid - l + 1
			//其实相当于从最底层也就是只有两个数的时候
			//加了一个判断，即使后面交换了也不影响
			//不如说交换后，会剔除重复的逆序对
		}
	}
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
	for i, j := left, 0; i <= right; i, j = i+1, j+1 {
        q[i] = tmp[j] // 把临时有序序列存回q序列中
    }

	return res
}