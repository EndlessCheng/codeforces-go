package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF598D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, m, q, x, y int
	Fscan(in, &n, &m, &q)
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for ; q > 0; q-- {
		Fscan(in, &x, &y)
		x--
		y--
		if ans[x][y] == 0 {
			todo, c := []pair{}, 0
			ans[x][y] = 1
			for q := []pair{{x, y}}; len(q) > 0; {
				p := q[0]
				q = q[1:]
				todo = append(todo, p)
				for _, d := range dir4 {
					if x, y := p.x+d.x, p.y+d.y; g[x][y] == '*' {
						c++
					} else if ans[x][y] == 0 {
						ans[x][y] = 1
						q = append(q, pair{x, y})
					}
				}
			}
			for _, p := range todo {
				ans[p.x][p.y] = c
			}
		}
		Fprintln(out, ans[x][y])
	}
}

//func main() { CF598D(os.Stdin, os.Stdout) }
