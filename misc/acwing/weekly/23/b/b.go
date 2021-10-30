package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x0, y0, x1, y1 int
	Fscan(in, &n, &x0, &y0, &x1, &y1)
	g := make([]string, n)
	for i := range g {
		Fscan(in, &g[i])
	}

	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	do := func(x, y int) [][]bool {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		var f func(int, int)
		f = func(x, y int) {
			if vis[x][y] {
				return
			}
			vis[x][y] = true
			for _, d := range dir4 {
				if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < n && g[xx][yy] == '0' {
					f(xx, yy)
				}
			}
		}
		f(x-1, y-1)
		return vis
	}
	vis0 := do(x0, y0)
	vis1 := do(x1, y1)

	ans := int(1e9)
	for i, r := range vis0 {
		for j, v := range r {
			if v {
				for ii, r := range vis1 {
					for jj, v := range r {
						if v {
							ans = min(ans, (i-ii)*(i-ii)+(j-jj)*(j-jj))
						}
					}
				}
			}
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
