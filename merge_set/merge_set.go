package main

import (
    "fmt"
    "os"
    "bufio"
    )

const N int = 100010
var p = make([]int, N)

//查询x祖宗节点，并路径压缩
func find(x int) int {
    if p[x] != x {
        p[x] = find(p[x])   //每次如果x所在集合不是独立存在，也就是不指向它本身
                            //那就递归找它指向的父节点的节点，它的父节点如果指向了自己，那它就是祖宗
    }
    return p[x]
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    var n, m int
    fmt.Fscan(reader, &n, &m)
    var (
        op string
        a int
        b int
    )
    
    //类似求前缀和，在外部由for去定义一个函数？
    for i := 1; i <= n; i++ {
        p[i] = i    //构建初始n个集合，他们的父节点就是他们本身
    }
    
    for i := 0; i < m; i++ {
        fmt.Fscan(reader, &op, &a, &b)
        if op == "M" {
            p[find(a)] = find(b)    //这一步的意思是，将a插入到b中
                                    //相当于a认了b为祖宗
        } else {
            if find(a) == find(b) {
                fmt.Println("Yes")
            } else {
                fmt.Println("No")
            }
        }
    }
}