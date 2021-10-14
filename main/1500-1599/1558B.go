package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1558B(in io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m int
	Fscan(in, &n, &m)
	f := make([]int, n+2)
	f[n] = 1
	for i := n - 1; i > 0; i-- {
		f[i] = f[i+1] * 2 % m
		for j := 2; j*i <= n; j++ {
			f[i] = (f[i] + (f[i*j]-f[min((i+1)*j, n+1)]+m)%m) % m
		}
	}
	Fprint(out, (f[1]-f[2]+m)%m)
}

//func main() { CF1558B(os.Stdin, os.Stdout) }
