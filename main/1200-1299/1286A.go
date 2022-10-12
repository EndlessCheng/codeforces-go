package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1286A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var n, v int
	Fscan(in, &n)
	f := make([][2]int, n/2+1)
	for i := 1; i <= n/2; i++ {
		f[i] = [2]int{1e9, 1e9}
	}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for j := n / 2; j >= 0; j-- {
			if v == 0 || v%2 > 0 {
				f[j][1] = min(f[j][1], f[j][0]+1)
			} else {
				f[j][1] = 1e9
			}
			if j > 0 && v%2 == 0 {
				f[j][0] = min(f[j-1][0], f[j-1][1]+1)
			} else {
				f[j][0] = 1e9
			}
		}
	}
	Fprint(out, min(f[n/2][0], f[n/2][1]))
}

//func main() { CF1286A(os.Stdin, os.Stdout) }
