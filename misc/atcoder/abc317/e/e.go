package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, sx, sy, tx, ty int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
		for j, b := range a[i] {
			if b == 'S' {
				sx, sy = i, j
			} else if b == 'G' {
				tx, ty = i, j
			}
		}
	}
	_ = tx+ty

	for _, row := range a {
		for i, b := range row {
			if b == '>' {
				for j := i + 1; j < m; j++ {
					if row[j] == '.' {
						row[j] = 'a'
					} else {
						break
					}
				}
			}
		}
		for i := len(row) - 1; i >= 0; i-- {
			b := row[i]
			if b == '<' {
				for j := i - 1; j >= 0; j-- {
					if row[j] == '.' {
						row[j] = 'a'
					} else {
						break
					}
				}
			}
		}
	}

	for j := range a[0] {
		for i, r := range a {
			b := r[j]
			if b == 'v' {
				for k := i + 1; k < n; k++ {
					if a[k][j] == '.' || a[k][j] == 'a' {
						a[k][j] = 'a'
					} else {
						break
					}
				}
			}
		}
		for i := len(a) - 1; i >= 0; i-- {
			b := a[i][j]
			if b == '^' {
				for k := i - 1; k >=0; k-- {
					if a[k][j] == '.' || a[k][j] == 'a' {
						a[k][j] = 'a'
					} else {
						break
					}
				}
			}
		}
	}

	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = -1
		}
	}
	dis[sx][sy] = 0
	type pair struct{ x, y int }
	q := []pair{{sx, sy}}
	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dir4 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < m && dis[x][y] < 0 && (a[x][y] == '.' || a[x][y] == 'G') {
					if a[x][y] == 'G' {
						Fprintln(out, step)
						return
					}
					dis[x][y] = step
					q = append(q, pair{x, y})
				}
			}
		}
	}
	Fprintln(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
