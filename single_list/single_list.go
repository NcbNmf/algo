package main
import (
	"fmt"
)

const N int = 100010

var head, idx int
var e [N]int
var ne [N]int

func new() {
	head = -1 //也就是指向头节点，但是因为现在没插，是空节点
	idx = 0 //一个指针，指向的最新能够用的节点
	//虽然idx是自然增长的，相当于是数组的下边，但实际链表却不是按照这个顺序来的
	//所以这个idx也可以指插入了几个点
	//当idx=1时，插入了1个，那个插入点的下标为0
	//当idx=k时，插入了k个，那个插入点的下标为k-1
}

//将x插入到头节点，这里指的是插入到head指向的下一个
func add_to_head(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

//在第k个输入的数后面插入x
//
func add(k, x int) {
	e[idx] = x		//相当于先创建一个节点，把数储存起来嘛
	ne[idx] = ne[k]	//然后这个节点的指针指向原本的k指向的值
	ne[k] = idx		//最后再让k指向这个节点就好
	idx++			//随后指针向后一位
}

//删除第k个输入的数后面的数，当k=0时代表删除头节点
func remove(k int) {
	ne[k] = ne[ne[k]] 
}

func traverse(head int) []int {
	res := []int{}
	for i := head; i != -1; i = ne[i] {	//这里的遍历就是让i不断的成为指针指向下一个值
		res = append(res, e[i])
	}
	return res
}

func main() {
	var m int
	fmt.Scanf("%d", m)

	new()

	for i := 0; i < m; i++ {
		var op string
		fmt.Scanf("%s", &op)
		switch op {
		case "H":
			var x int
			fmt.Scanf("%d", &x)
			add_to_head(x)
		case "I":
			var k, x int
			fmt.Scanf("%d%d", &k, &x)
			add(k-1, x)		//这里为什么是k-1，因为，第k个输入的数的下标就是k-1
		case "D":
			var k int
			fmt.Scanf("%d", &k)
			if k == 0 {
				head = ne[head] //这里删除头节点，意思是让head连接到原本head往后的第二个节点
			} else {
				remove(k-1)
			}
		}
	}

	res := traverse(head)
	for i := range res {
		fmt.Printf("%d ", res[i])
	}
}

