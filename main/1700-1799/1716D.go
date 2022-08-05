package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1716D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var n, k int
	Fscan(in, &n, &k)
	ans := make([]int, n+1)
	f := make([]int, n+1)
	f[0] = 1
	for i, s := k, 0; s+i <= n; i++ {
		s += i
		g := make([]int, n+1)
		for j := i; j <= n; j++ {
			g[j] = (g[j-i] + f[j-i]) % mod
			ans[j] = (ans[j] + g[j]) % mod
		}
		f = g
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF1716D(os.Stdin, os.Stdout) }
