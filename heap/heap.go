package main

import (
	"fmt"
	"os"
	"bufio"
)

const N int = 100010

var n, size, idx int
var (
	heap [N]int			//储存堆中的值
	ph [N]int	//储存第k个插入的点在堆中的位置，也就是下标
	hp [N]int	//储存堆中下标是k的点是第几个插入的
)

func swap(a, b int) {
	ph[hp[a]], ph[hp[b]] = ph[hp[b]], ph[hp[a]]
    hp[a], hp[b] = hp[b], hp[a]
    heap[a], heap[b] = heap[b], heap[a]
}
/*
每次维护堆的时候，插入、删除、修改都会影响点位置，因此需要通过swap来映射一个抽象的堆
在实际操作中
如果我们只有插入操作（从队尾插入），那么下标size和节点idx是完全对应的
但当节点位置发生变化，idx与size的关系会被打破（size变化）
因此需要一个数组来储存每个size下标对应的idx的值，也就是hepa_point[size]=idx
同时也需要储存第idx个插入的数的size下标，也就是point_heap[idx]=size
当然了也需要有一个heap[]=x来记录值

那么swap怎么解释呢
如：删除一个数的操作，我们通常是把它和队尾的数进行交换然后size--
这个时候原本位置的那个节点它的下标size没变，但是idx第几次插入却是变了的
更别提之后down操作重新排序
因此size需要重新映射（指向）
先将他们的下标，也就是ph[]交换
然后hp[]也就是第几次插入的标记交换
然后再是值h[]交换

*/
func up(x int) {
	for x/2 > 0 && heap[x/2] > heap[x] {//判断存在父节点，并且父节点更大
		swap(x/2, x)
		x /= 2	//这里就是让父亲变儿子
	}
}

func down(x int) {
	tmp := x
	if x * 2 <= size && heap[x*2] < heap[tmp] {	
		tmp = x * 2							
	}
	if x * 2 + 1 <= size && heap[x*2+1] < heap[tmp] {
		tmp = x * 2 + 1
	}
	if x != tmp {	
		swap(x, tmp)
		down(tmp)	
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	var op string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &op)
		switch op {
		case "I":
			var v int
			fmt.Fscan(in, &v)
			idx++
			size++
			ph[idx] = size
			hp[size] = idx
			heap[size] = v
			up(size)
		case "PM":
			fmt.Println(heap[1])
		case "DM":
			swap(1, size)
			size--
			down(1)
		case "D":
			var k int
			fmt.Fscan(in, &k)
			k = ph[k]	//先将第k个插入的数的下标赋予k
						//为什么赋下标，因为所有的up和down都是基于下标实现的
			swap(k, size)
			size--
			down(k)
			up(k)	//这里直接来一轮up和down是因为无论变大变小
					//接触第一层for后就会退出，只执行一个
		case "C":
			var k, v int
			fmt.Fscan(in, &k, &v)
			k = ph[k]
			heap[k] = v
			down(k)
			up(k)
		}
	}
}