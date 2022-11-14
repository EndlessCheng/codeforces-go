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
	var n, m int
	Fscan(in, &n, &m)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		f[i][0] = i
	}
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &b[i])
		f[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i] == b[j] {
				f[i][j] = min(min(f[i-1][j], f[i][j-1])+1, f[i-1][j-1])
			} else {
				f[i][j] = min(min(f[i-1][j], f[i][j-1])+1, f[i-1][j-1]+1)
			}
		}
	}
	Fprint(out, f[n][m])
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
