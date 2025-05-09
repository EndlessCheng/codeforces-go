package main

import (
	"bufio"
	. "fmt"
	"io"
	"runtime/debug"
)

// https://github.com/EndlessCheng
func cf525D(in io.Reader, _w io.Writer) {
	debug.SetMemoryLimit(500 << 20)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m uint16
	Fscan(in, &n, &m)
	g := make([][]byte, n+2)
	g[0] = make([]byte, m+2)
	for i := uint16(1); i <= n; i++ {
		var s []byte
		Fscan(in, &s)
		g[i] = make([]byte, m+2)
		copy(g[i][1:], s)
	}
	g[n+1] = make([]byte, m+2)

	var dfs func(uint16, uint16)
	dfs = func(x, y uint16) {
		if g[x][y] == '.' ||
			!(g[x-1][y-1] == '.' && g[x-1][y] == '.' && g[x][y-1] == '.' ||
				g[x-1][y+1] == '.' && g[x-1][y] == '.' && g[x][y+1] == '.' ||
				g[x+1][y-1] == '.' && g[x+1][y] == '.' && g[x][y-1] == '.' ||
				g[x+1][y+1] == '.' && g[x+1][y] == '.' && g[x][y+1] == '.') {
			return
		}
		g[x][y] = '.'
		for i := max(x-1, 1); i <= min(x+1, n); i++ {
			for j := max(y-1, 1); j <= min(y+1, m); j++ {
				if i != x || j != y {
					dfs(i, j)
				}
			}
		}
	}
	for i := uint16(1); i <= n; i++ {
		for j := uint16(1); j <= m; j++ {
			dfs(i, j)
		}
	}

	for _, row := range g[1 : n+1] {
		Fprintf(out, "%s\n", row[1:m+1])
	}
}

//func main() { cf525D(bufio.NewReader(os.Stdin), os.Stdout) }
