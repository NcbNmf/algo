package main 

import (
	"fmt"
	"os"
	"bufio"

)

const N int = 1000010

var (
	a, q [N]int
	in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	//in就是读取位置
	//out就是输出位置
)

func main() {
	defer out.Flush()
	//该方法可以将缓冲中的数据写入下层io.Writer
	//也就是和bufio.NewWriter搭配使用

    var n, k int
    fmt.Fscan(in, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
    
	hh, tt := 0, -1
	for i := 0; i < n; i++ {
		//首先判断队头是否已经滑出窗口
		//i-k+1就是队列头
		if hh <= tt && i-k+1 > q[hh] {
			hh++
		}
		//接着构建递增队列
		//求最小值
		for hh <= tt && a[q[tt]] >= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		//下面的判断是为了确定窗口已完全进入数组
		//如窗口k=3，如果判断下标没有大于k-1=2，那就是窗口没有完全进入
		//不进行最大最小值输出
		if i >= k-1 {
			fmt.Fprintf(out, "%d ", a[q[hh]])
		}
	}
	fmt.Fprintf(out, "\n")

	//初始化队头队尾指针，进行另一个最值队列构建
	hh, tt = 0, -1
	for i := 0; i < n; i++ {
		if hh <= tt && i-k+1 > q[hh] {
			hh++
		}
		//接着构建递减队列
		//求最大值
		for hh <= tt && a[q[tt]] <= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Fprintf(out, "%d ", a[q[hh]])
		}
	}
	fmt.Fprintf(out, "\n")

}