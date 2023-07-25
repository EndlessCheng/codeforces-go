package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF193A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, t, cnt, tot int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += bytes.Count(a[i], []byte("#"))
	}
	if tot < 3 {
		Fprint(out, -1)
		return
	}

	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	vis := make([][]int, n)
	for i := range vis {
		vis[i] = make([]int, m)
	}
	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] == t || a[x][y] != '#' {
			return false
		}
		vis[x][y] = t
		cnt++
		for _, d := range dir4 {
			dfs(x+d.x, y+d.y)
		}
		return true
	}
	for i, row := range a {
		for j, b := range row {
			if b != '#' {
				continue
			}
			row[j] = '.'
			t++
			cnt = 0
			for _, d := range dir4 {
				if dfs(i+d.x, j+d.y) {
					break
				}
			}
			if cnt < tot-1 {
				Fprint(out, 1)
				return
			}
			row[j] = '#'
		}
	}
	Fprint(out, 2)
}

//func main() { CF193A(os.Stdin, os.Stdout) }
