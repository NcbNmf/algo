package main 
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &m)
	var stack = make([]int, 0)

	for m > 0 {
		m--
		var op string
		fmt.Fscan(reader, &op)
		switch op {
		case "push":
			var x int
			fmt.Fscan(reader, &x)
			stack = append(stack, x)
		case "pop":
			stack = stack[:len(stack)-1]	//因为后面的是开区间
											//所以可以模拟弹出操作
		case "empty":
			if len(stack) == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "query":
			fmt.Println(stack[len(stack)-1])
		}
	}
}