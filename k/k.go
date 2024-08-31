package main
import "fmt"

func main() {

	var (
		n int
		k int
	)
	fmt.Scanf("%d %d", &n, &k)	//这里注意不要加逗号，用空格，原理未知
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	fmt.Pritln(quick_K(0, n - 1, k, q))
}

func quick_K(left, right, k int, q []int) int {

	if left == right {
		return q[left]
	}

	x := q[(left+right)/2]
	l, r := left-1, right+1
	
	for l < r {
		for {
			l++
			if q[l] >= x {
				break
			}
		}
		for {
			r--
			if q[r] <= x {
				break
			}
		}
		if l < r {
			q[l], q[r] = q[r], q[l]
		}
	}

	xx := r - left + 1 
	//核心行
	//表示了分好后，左边有多少个数
	if k <= xx {
		return quick_K(left, r, k, q)
	}
	return quick_K(r+1, right, k-xx, q)
	//这里k-xx 是指第k个数再分好后的右边的第几个
	//一直递归到left=right，就得出第k的数的大小了
	

}