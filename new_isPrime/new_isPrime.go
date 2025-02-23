package main

import (
    "fmt"
    "os"
    "bufio"
    )
    
func main() {
    in := bufio.NewReader(os.Stdin)
    
    var n, x int
    fmt.Fscan(in, &n)
    
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &x)
        if isPrime(x) {
            fmt.Println("Yes")
        } else {
            fmt.Println("No")
        }
    }
    
    
}

func isPrime(x int) bool {
    if x == 1 {
        return false
    }
    for i := 2; i <= x/i; i++ {	//试除法
        if x % i == 0 {
            return false
        }
    }
    return true
}