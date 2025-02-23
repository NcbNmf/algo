package main
import (
	"fmt"
	"bufio"
	"os"
)

const N int = 100010

var (
	son [][]int	//节点
	cnt []int	//单词尾部计数
	idx int		//为节点作标记
)

func insert(str string) {
	p := 0 //首先从根节点开始，结合图去推会清楚一点
	for i := 0; i < len(str); i++ {
		u := str[i]-'a' //这是让字符映射成ascll码值的输入方式
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx	//idx的作用是为每个节点添加一个标记，独一无二
							//做标记还有一个作用，创建节点，有值了他才能算是存在
		}
		p = son[p][u] //这时候更新节点为止
	}
	cnt[p]++	//并且在for结束后，为在这个节点结尾的字符串做标记+1
				//这里怎么判断呢，因为idx给每个节点做了标记
				//因此每个节点p都是是特殊的
}

func query(str string) {
	p := 0
	for i := 0; i < len(str); i++ {
		u := str[i]-'a'
		if son[p][u] == 0 {
			fmt.Println(0)	//这里直接输出不存在是因为
							//查询枚举过程中，只要有一个节点不符合
							//那这个单词就是不存在	
			return
		}
		p = son[p][u]	//别忘了指向下一个节点
	}
	fmt.Println(cnt[p])
	return
}

func main() {
	//因为二维切片不好在main外设置大小
	//因此需要先把切片的大小创建
	for i := 0; i < N; i++ {
		tmp := make([]int, 26)
		son = append(son, tmp)
	}
	cnt = make([]int, N)

	reader := bufio.NewReader(os.Stdin)
	
	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var op string
		var x string
		fmt.Fscan(reader, &op, &x)
		if op == "I" {
			insert(x)
		} else if op == "Q" {
			query(x)
		}
	}
}

