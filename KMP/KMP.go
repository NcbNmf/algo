package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var p, s int
	var patt, sum []byte	//string输入[]byte会自动变成byte值
	fmt.Fscan(reader, &p, &patt, &s, &sum)
    next := make([]int, p)

    for i, j := 1, 0; i < p; i++ {
        for j > 0 && patt[i] != patt[j] {
            j = next[j-1]
        }
        if patt[i] == patt[j] {
            j++
        }
        next[i] = j
    }
	for i, j := 0, 0; i < s; i++ {
        for j > 0 && (j == p || sum[i] != patt[j]) {
            j = next[j-1]
        }
        if sum[i] == patt[j] {
            j++
        }
        if j == p {
            fmt.Fprintf(writer ,"%d ", i-p+1)
        }
    }


}