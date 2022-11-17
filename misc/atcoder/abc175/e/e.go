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
	var n, m, k, x, y, v int
	Fscan(in, &n, &m, &k)
	a := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, m+1)
	}
	for ; k > 0; k-- {
		Fscan(in, &x, &y, &v)
		a[x][y] = v
	}
	f := make([][][4]int, n+1)
	for i := range f {
		f[i] = make([][4]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j][0] = max(f[i][j-1][0], max(max(f[i-1][j][1], f[i-1][j][2]), f[i-1][j][3]))
			f[i][j][1] = max(f[i][j-1][1], f[i][j][0]+a[i][j])
			f[i][j][2] = max(f[i][j-1][2], f[i][j-1][1]+a[i][j])
			f[i][j][3] = max(f[i][j-1][3], f[i][j-1][2]+a[i][j])
		}
	}
	Fprint(out, max(max(f[n][m][1], f[n][m][2]), f[n][m][3]))
}

func main() { run(os.Stdin, os.Stdout) }
func max(a, b int) int { if b > a { return b }; return a}
