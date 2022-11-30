package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m int
	Fscan(in, &n, &m)
	c := make([]int, n+1)
	c[0] = 1
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			c[j] = (c[j] + c[j-1]) % mod
		}
	}
	f := make([]int, m+1)
	f[0] = 1
	for i := 2; i <= m; i += 2 {
		for j := 0; j <= i && j <= n; j += 2 {
			f[i] = (f[i] + f[(i-j)/2]*c[j]) % mod
		}
	}
	Fprint(out, f[m])
}

func main() { run(os.Stdin, os.Stdout) }
