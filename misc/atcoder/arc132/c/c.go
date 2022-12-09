package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	var n, d, v, ans int
	Fscan(in, &n, &d)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 1<<(d*2))
	}
	f[0][0] = 1
	for i, fs := range f[:n] {
		Fscan(in, &v)
		for mask, res := range fs {
			for j := max(i-d, 0); j <= min(i+d, n-1); j++ {
				if (v == -1 || j == v-1) && mask>>(j-i+d)&1 == 0 {
					m := (mask | 1<<(j-i+d)) >> 1
					f[i+1][m] = (f[i+1][m] + res) % mod
				}
			}
		}
	}
	for _, res := range f[n] {
		ans += res
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
