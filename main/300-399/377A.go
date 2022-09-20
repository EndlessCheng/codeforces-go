package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF377A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, m, k int
	Fscan(in, &n, &m, &k)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var dfs func(int, int)
	dfs = func(i, j int) {
		vis[i][j] = true
		for _, d := range dir4 {
			if x, y := i+d.x, j+d.y; 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && a[x][y] == '.' {
				dfs(x, y)
			}
		}
		if k > 0 {
			k--
			a[i][j] = 'X'
		}
	}
	for i, row := range a {
		for j, v := range row {
			if v == '.' {
				dfs(i, j)
				for _, r := range a {
					Fprintf(out, "%s\n", r)
				}
				return
			}
		}
	}
}

//func main() { CF377A(os.Stdin, os.Stdout) }
