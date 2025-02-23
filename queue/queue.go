package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m int
	fmt.Fscan(reader, &m)
	var queue = make([]int, 0)
	var op string
	for m > 0 {
		m--
		fmt.Fscan(reader, &op)
		switch op {
		case "push":
			var x int
			fmt.Fscan(reader, &x)
			queue = append(queue, x)
		case "pop":
			queue = queue[1:]	//相当于弹出第一个元素
								//用切片头当作队头
		case "empty":
			if len(queue) == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "query":
			fmt.Println(queue[0])
		}
	}
}