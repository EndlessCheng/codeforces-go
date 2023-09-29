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
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	f := make([][3001]int, n)
	for j := a[0]; j <= b[0]; j++ {
		f[0][j] = j - a[0] + 1
	}
	for i := 1; i < n; i++ {
		for j := a[i]; j <= b[i]; j++ {
			// max(j-1, 0) 是防止 j=0 时产生 -1 下标
			f[i][j] = (f[i][max(j-1, 0)] + f[i-1][min(j, b[i-1])]) % mod
		}
	}
	Fprint(out, f[n-1][b[n-1]])
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
