package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF540C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var n, m, sx, sy, tx, ty, c int
	Fscan(in, &n, &m)
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	Fscan(in, &sx, &sy, &tx, &ty)
	sx--
	sy--
	tx--
	ty--
	if sx == tx && sy == ty {
		for _, d := range dir4 {
			if x, y := tx+d.x, ty+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' {
				Fprint(out, "YES")
				return
			}
		}
		Fprint(out, "NO")
		return
	}
	g[sx][sy] = '.'
	if g[tx][ty] == '.' {
		for _, d := range dir4 {
			if x, y := tx+d.x, ty+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] == '.' {
				c++
			}
		}
		if c < 2 {
			Fprint(out, "NO")
			return
		}
	}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var f func(int, int) bool
	f = func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] {
			return false
		}
		if x == tx && y == ty {
			return true
		}
		if g[x][y] == 'X' {
			return false
		}
		vis[x][y] = true
		for _, d := range dir4 {
			if f(x+d.x, y+d.y) {
				return true
			}
		}
		return false
	}
	if f(sx, sy) {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF540C(os.Stdin, os.Stdout) }
