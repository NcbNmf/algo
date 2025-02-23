package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {

	var (
		n int
		m int
	)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m,)
	
	var A = make([]int, n)
	var B = make([]int, m)



	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &A[i])
	}
	for j := 0; j < m; j++ {
		fmt.Fscan(reader, &B[j])
	}
    
    i := 0
	for j := 0; i < n && j < m; j++ {	//注意这里不仅i<n，也要j<m，要不会超出范围
		if A[i] == B[j] {
			i++
		}
	}

	if i == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}