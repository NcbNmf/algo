package main

import (
	"fmt"
	"os"
	"bufio"
)

const N int = 100010

var n, m, size int
var heap [N]int

func down(x int) {
	//tmp存在的意义
	//意味着最小值标志
	tmp := x
	if x * 2 <= size && heap[x*2] < heap[tmp] {	//这两个条件分别是
											//判断左儿子是否存在
											//以及x节点的值是否小于左儿子
		tmp = x * 2							//如果小于，将左儿子标记为最小值
											//然后进入到下一个if进行比对
	}
	if x * 2 + 1 <= size && heap[x*2+1] < heap[tmp] {
		tmp = x * 2 + 1
	}
	if x != tmp {	//最后，如果最小值tmp不等于原来根节点就交换
		heap[x], heap[tmp] = heap[tmp], heap[x]
		down(tmp)	//再最后，因为已经经过了交换，也就是tmp指向的不是最小值了
					//而是指向交换后的较大的儿子
					//那另一个儿子怎么办？
					//照理来说从n往前--遍历，每一个都能down到
					//但是从n/2往前也可以，因为二叉树最后一层就是n/2个元素
	}

}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in,  &heap[i])
	}
	size = n
	//直接构建一位数组，

	//然后就是把一位数组变成堆
	//如果从n开始down，那么时间复杂度会很高
	//但是从n/2的节点开始down就会变成O(n)
	//并且因为二叉树的最后一层元素右n/2个，所以无论如何都能遍历完全
	for i := n/2; i != 0; i-- {
		down(i)
	}

	for i := 0; i < m; i++{
		fmt.Printf("%d ", heap[1])	//每次输出堆顶
		heap[1] = heap[size]			//再令堆顶等于最后一个数
		size--						//去除最小值
		down(1)						//再重新排序
	}
}