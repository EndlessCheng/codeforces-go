package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1840F(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 1<<12)
	_i := len(buf)
	rc := func() byte {
		if _i >= len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	ri := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	type tuple struct{ x, y, c int }
	dir3 := []tuple{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

o:
	for T := ri(); T > 0; T-- {
		n, m, r := ri(), ri(), ri()
		type pair struct{ i, t int }
		die := [2]map[pair]bool{{}, {}}
		for i := 0; i < r; i++ {
			t, tp, c := ri(), ri(), ri()
			die[tp-1][pair{c, t}] = true
		}
		vis := make([][][]bool, n+1)
		for i := range vis {
			vis[i] = make([][]bool, m+1)
			for j := range vis[i] {
				vis[i][j] = make([]bool, r+1)
			}
		}
		vis[0][0][0] = true
		q := []tuple{{}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range dir3 {
				x, y, c := p.x+d.x, p.y+d.y, p.c+d.c
				if x <= n && y <= m && c <= r && !vis[x][y][c] && !die[0][pair{x, x + y + c}] && !die[1][pair{y, x + y + c}] {
					if x == n && y == m {
						Fprintln(out, x+y+c)
						continue o
					}
					vis[x][y][c] = true
					q = append(q, tuple{x, y, c})
				}
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { CF1840F(os.Stdin, os.Stdout) }
