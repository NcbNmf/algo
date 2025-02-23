package main

import (
	"fmt"
	"os"
	"bufio"
)

const N = 300010

func quickSort(q []int, left, right int) {
	if left >= right {
		return
	}
	l, r := left-1, right+1
	x := q[l+(r-l) >> 1]
	
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

	quickSort(q, left, r)
	quickSort(q, r+1, right)
}

func findx(q[]int, x int) int {
	l := 0
	r := len(q) - 1
	for l < r {
		mid := (l+r) >> 1
		if q[mid] >= x {
			r = mid
		} else {
			l = mid + 1	//这里加1是因为mid=(l+r)/2是向下取整
						//逼近到边界时，如果一直是l=mid就跳不出循环
		}
	}
	return l+1 //不明白为什么+1
				//可能是因为去重后下标是从1开始
				//即离散化后的坐标一般从1开始
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	var x,c,l,r int
	step := make([]int, 0)
	addTable := make([][]int, 0)
	queryTable := make([][]int, 0)
	var q [N]int
	var a [N]int
	

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &x, &c)
		addTable = append(addTable, []int{x, c})
		step = append(step, x)
	}

	for i := 1; i <= m; i++ {
		fmt.Fscan(reader, &l, &r)
		step = append(step, []int{l,r}...)	//想一次插两个元素进去时用切片解包的方法
		queryTable = append(queryTable, []int{l,r})
	}

	//上面的就是离散化的第一步，下一步需要把下标和它实际指向的数联系起来
	//排序后去重
	quickSort(step, 0, len(step) - 1)

	j := 1
	for i := 1; i < len(step); i++ {
		if step[i] != step[i-1] {
			step[j] = step[i]
			j++
		}
	}
	step = step[0:j]

	for i := 0; i < len(addTable); i++ {
		index := findx(step, addTable[i][0])
		q[index] += addTable[i][1]	//这里是二维切片
	}

	for i := 1; i <= len(step); i++ {
		a[i] = a[i-1] + q[i]
	}

	for i := 0; i < len(queryTable); i++ {
		left := findx(step, queryTable[i][0])
		right := findx(step, queryTable[i][1])
		fmt.Println(a[right] - a[left-1])
	}
}