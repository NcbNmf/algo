package main 

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)



// func quickSort(q [][]int, left, right int) {
// 	if left > right {
// 		return
// 	}
// 	mid := q[(right + left)/2][0]
// 	l := left - 1
// 	r := right + 1
	
// 	for l < r {
// 		for {
// 			l++
// 			if q[l][0] >= mid {
// 				break
// 			}
// 		}
// 		for {
// 			r--
// 			if q[r][0] <= mid {
// 				break
// 			}
// 		}
// 		if l < r {
// 			q[r][0], q[l][0] = q[l][0], q[r][0]
// 		}
// 	}

// 	quickSort(q, left, r)
// 	quickSort(q, r+1, right)
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	var interval = make([][]int, n)
	var l, r int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l, &r)
		interval[i] = []int{l, r}
	}

	//排序
	//快排占用内存太大，容易MLT
	// quickSort(interval, 0, len(interval)- 1)
	sort.Slice(interval, func(i, j int) bool {
        return interval[i][0] < interval[j][0] ||
            (interval[i][0] == interval[j][0] && interval[i][1] > interval[j][1])
    })

	var res = make([]int, 2)
	res[0] = interval[0][0]
	res[1] = interval[0][1]
	var j int
	for i := 0; i < n; i++ {
		if i == n-1 && interval[i][0] <= res[1] {
			j++
		}
		if i == n-1 && interval[i][0] > res[1] {
			j += 2
		}
		
		if i != n-1 && res[1] < interval[i+1][0] {
			j++
			res[0] = interval[i+1][0]
			res[1] = interval[i+1][1]

		}
		if i != n-1 && res[1] >= interval[i+1][0] {
			res[1] = max(interval[i+1][1],res[1])
		} 
	}

	//上面是大白话，下面是诗，PS:太漂亮了
	//第一个if，可以看作是更新左端点
	//else后面看作更新右端点
	// var res [][]int
    // res = append(res, intervals[0])
    // for _, interval := range intervals[1:] {
    //     if interval[0] > res[len(res)-1][1] {
    //         res = append(res, interval)
    //     } else {
    //         res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
    //     }
    // }
    // fmt.Println(len(res))


	fmt.Println(j)
}