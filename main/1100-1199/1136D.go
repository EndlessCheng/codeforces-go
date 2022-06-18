package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1136D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, x, y int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	swap := make(map[[2]int]bool, m)
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		swap[[2]int{x, y}] = true
	}
	mustSwap := a[n-1:]
	for i := n - 2; i >= 0; i-- {
		x := a[i]
		for _, y := range mustSwap {
			if !swap[[2]int{x, y}] {
				mustSwap = append(mustSwap, x)
				break
			}
		}
	}
	Fprint(out, n-len(mustSwap))
}

//func main() { CF1136D(os.Stdin, os.Stdout) }
