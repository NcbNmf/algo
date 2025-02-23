//因为这个题前两个要求和合并集合的非常像
//因此代码可以直接复制

package main

import (
    "fmt"
    "os"
    "bufio"
    )

const N int = 100010
var p = make([]int, N)
var size = make([]int, N)   

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
    //相当于初始化p[]和size[]
    for i := 1; i <= n; i++ {
        p[i] = i    //构建初始n个集合，他们的父节点就是他们本身
        size[i] = 1 //n个集合，初始化每个集合只有它自身，也就是个数是1
    }
    
    for i := 0; i < m; i++ {
        fmt.Fscan(reader, &op)
        
        if op == "C" {
            fmt.Fscan(reader, &a, &b)
            if find(a) == find(b) {
                continue            //这里需要特判，如果ab相等，就不需要连通了
            }
            size[find(b)] += size[find(a)]  //注意是b+=a
                                            //这一步巧妙在于，一开始只有两个节点连通那就是1+1
                                            //后面再不断+=
            p[find(a)] = find(b)    //这一步的意思是，将a插入到b中
                                    //相当于a认了b为祖宗
                                    //在连通块中，也是相当于把a和b连通
                                    
        } else if op == "Q1" {
            fmt.Fscan(reader, &a, &b)
            if find(a) == find(b) {
                fmt.Println("Yes")
            } else {
                fmt.Println("No")
            }
            
        } else {
            fmt.Fscan(reader, &a)
            fmt.Println(size[find(a)])
        }
    }
}