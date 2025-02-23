package main
import (
	"fmt"
)

const N int = 100010

var lne [N]int
var rne [N]int
var e [N]int
var idx int

//初始化双链表，也就是数组下标的0和1成为两端
//然后让他们连个互相指向就可以了
func new() {
	rne[0] = 1
	lne[1] = 0
	idx = 2	//因为数组前两个位置都被占了，所以idx从2开始
}

func add(k, x int) {
	//和单链表类似，需要把需要插入的节点构建好
	//也就是储存值，再加插入节点指向两边
	e[idx] = x
	rne[idx] = rne[k]
	lne[idx] = k
	//然后这里需要先构建好原本k右边节点指向idx的指针
	//因为后面rne[k]需要指向idx节点，会变化
	lne[rne[k]] = idx
	rne[k] = idx
	idx++ 
}

func remove(k int) {
	//只需要将k的左右两个点链接起来就好
	rne[lne[k]] = rne[k]
	lne[rne[k]] = lne[k]
}

func traverse(head int) []int {
	res := []int{}
	for i := head; i != 1; i = rne[i] {	//这里的遍历就是让i不断的成为指针指向下一个值
		res = append(res, e[i])
	}
	return res
}

func main() {
	var m int
	fmt.Scanf("%d", &m)

	new()	

	for i := 0; i < m; i++ {
		var op string
		fmt.Scanf("%s", &op)
		switch op {
		case "L":
			var x int
			fmt.Scanf("%d", &x)
			add(0, x)
		case "R":
			var x int
			fmt.Scanf("%d", &x)
			add(lne[1], x)
		case "D":
			var k int
			fmt.Scanf("%d", &k)
			remove(k+1)	//这里k+1是因为idx从2开始，也就是第一个插入的数下标为2
		case "IL":
			var k, x int
			fmt.Scanf("%d%d", &k, &x)
			add(lne[k+1], x)
		case "IR":
			var k, x int
			fmt.Scanf("%d%d", &k, &x)
			add(k+1, x)
		}
	}
	res := traverse(rne[0])
	for i := range res {
		fmt.Printf("%d ", res[i])
	}
}