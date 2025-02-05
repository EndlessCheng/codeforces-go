package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 50
	dir4 := []struct {
		x, y int8
		c    byte
	}{{1, 0, 'D'}, {0, -1, 'L'}, {-1, 0, 'U'}, {0, 1, 'R'}}

	t0 := time.Now()

	var sx, sy int8
	var c, ans, sum int
	Fscan(in, &sx, &sy)
	g := [mx][mx]int16{}
	for i := range g {
		for j := range g[i] {
			Fscan(in, &g[i][j])
		}
	}
	a := [mx][mx]int8{}
	for i := range a {
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}

	st := []int8{}
	for i, d := range dir4 {
		if xx, yy := sx+d.x, sy+d.y; 0 <= xx && xx < mx && 0 <= yy && yy < mx && g[xx][yy] != g[sx][sy] {
			st = append(st, int8(i))
		}
	}
	var limitBase time.Duration
	if len(st) == 1 {
		limitBase = 1950 * time.Millisecond
	} else if len(st) == 2 {
		limitBase = 975 * time.Millisecond
	} else if len(st) == 3 {
		limitBase = 650 * time.Millisecond
	} else {
		limitBase = 488 * time.Millisecond
	}
	limit := limitBase

	path := []byte{}
	ansPath := ""
	vis := [mx * mx]bool{}
	var f func(x, y, pre int8) bool
	f = func(x, y, pre int8) bool {
		if c&(1<<16-1) == 0 && time.Since(t0) > limit {
			return true
		}
		c++
		sum += int(a[x][y])
		if sum > ans {
			ans = sum
			ansPath = string(path)
		}
		vis[g[x][y]] = true

		d := dir4[pre]
		if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < mx && 0 <= yy && yy < mx && !vis[g[xx][yy]] {
			path = append(path, d.c)
			if f(xx, yy, pre) {
				return true
			}
			path = path[:len(path)-1]
		}
		for i, d := range dir4 {
			if int8(i) == pre {
				continue
			}
			if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < mx && 0 <= yy && yy < mx && !vis[g[xx][yy]] {
				path = append(path, d.c)
				if f(xx, yy, int8(i)) {
					return true
				}
				path = path[:len(path)-1]
			}
		}

		vis[g[x][y]] = false
		sum -= int(a[x][y])
		return false
	}
	for _, i := range st {
		vis = [mx * mx]bool{}
		path = nil
		sum = 0
		f(sx, sy, i)
		limit += limitBase
	}
	Fprint(out, ansPath)
}

func main() { run(os.Stdin, os.Stdout) }
