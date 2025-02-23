package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)
	var s = make([]int, 100010)
	var a = make([]int, 100010)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	res := 0

	for i, j := 0, 0; i < n; i++ {
		s[a[i]]++
		for j < i && s[a[i]] > 1 {
			s[a[j]]--
			j++
		}
		if res < i - j + 1 {
			res = i - j + 1
		}
	}
	fmt.Println(res)
}