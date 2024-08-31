package main
import (
	"fmt"
)

func quickSort(q []int, left, right int) {
	if left >= right {
		return	//终止条件
	}
	x := q[(right + left)/2]
	//这里需要转一下，以保证后续这个left和right还能用
	l, r := left-1, right+1

	for l < r {	
		//退出条件
		//左指针
		for {
			l++
			if q[l] >= x {
				break
			}
			//边界问题，比如这里的判断，必须要有一个等于号
			//假如没有等于号，且当x位于两端时，甚至迭代时，x必然有在两端的那一刻
			//那当q[l]=x时无法break，++后循环就会出现下标超出
		}
		//右指针
		for {
			r--
			if q[r] <= x {
				break
			}
		}

		//这里加判断是为了确保前面--++后不会超过边界
		if l < r {
			q[l], q[r] = q[r], q[l]
		}
	}

	quickSort(q, left, r)
	quickSort(q, r+1, right)
	// quickSort(q, left, l-1)
	// quickSort(q, l, right)
	//这里必须对称，原理未知
	//不能
	// quickSort(q, left, r-1)
	// quickSort(q, r, right)
	
	//【0，0】l指向0，r指向1，这时候x=0，那进入到func后，l指向-1，r指向2，这时候还是会
	//进行判断，l还是指向0，r也指向1，进行交换也没用，这时候还没进入迭代，因为还是l<r
	//再一次循环后就变成了r指向0而l指向1，这时候再进入迭代，那就会变成
	//r=0=l，返回了上层迭代，而正确的迭代就不会有这种情况
	//本质上是让r不能接触到l，或是说让r不能成为l

	//边界问题解答
	//https://www.elefans.com/category/jswz/34/1261279.html
	//https://www.acwing.com/problem/content/discussion/content/3430/

}
func main() {

	var n int
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	quickSort(q, 0, n - 1)
	for i := 0; i < n; i++ {
        fmt.Printf("%d ", q[i])	//这是输出切片不带括号的方法
    }
    return
}

