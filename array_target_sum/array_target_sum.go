package main

import (
	"fmt"
	"os"
	"bufio"
)

const N = 100010	

func main() {

	var (
		n int
		m int
		x int
	)

	var A = make([]int, N)
	var B = make([]int, N)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &x)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &A[i])
	}
	for j := 0; j < m; j++ {
		fmt.Fscan(reader, &B[j])
	}
	
	j := m - 1	//从后往前遍历
	for i := 0; i < n; i++ {
		for j >= 0 && A[i]+B[j] > x {
			j--
		}
		if j >= 0 && A[i]+B[j] == x {
			fmt.Printf("%d %d", i, j)
		}
	}

}