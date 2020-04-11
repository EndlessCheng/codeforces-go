package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dir4 := [...][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	type pair struct{ x, y int }

	solveSet1 := func(_case int) (ans int) {
		var n, m int
		Fscan(in, &n, &m)
		mat := make([][]int, n)
		nei := make([][][4]pair, n)
		for i := range mat {
			mat[i] = make([]int, m)
			nei[i] = make([][4]pair, m)
			for j := range mat[i] {
				Fscan(in, &mat[i][j])
				for k, d := range dir4 {
					nei[i][j][k] = pair{i + d[0], j + d[1]}
				}
			}
		}
		for {
			removes := []pair{}
			for i, row := range mat {
				for j, v := range row {
					if v == 0 {
						continue
					}
					ans += v
					s, s2 := 0, 0
					for _, p := range nei[i][j] {
						if x, y := p.x, p.y; x >= 0 && x < n && y >= 0 && y < m && mat[x][y] > 0 {
							s += mat[x][y]
							s2 += v
						}
					}
					if s2 < s {
						removes = append(removes, pair{i, j})
					}
				}
			}
			if len(removes) == 0 {
				break
			}
			for _, p := range removes {
				mat[p.x][p.y] = 0
			}
			for i, row := range mat {
				for j, v := range row {
					if v == 0 {
						continue
					}
					for k, p := range nei[i][j] {
						if x, y := p.x, p.y; x >= 0 && x < n && y >= 0 && y < m && mat[x][y] == 0 {
							d := dir4[k]
							for {
								x += d[0]
								y += d[1]
								if !(x >= 0 && x < n && y >= 0 && y < m && mat[x][y] == 0) {
									nei[i][j][k] = pair{x, y}
									break
								}
							}
						}
					}
				}
			}
		}
		return
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: ", _case)
		Fprintln(out, solveSet1(_case))
	}
}

func main() { run(os.Stdin, os.Stdout) }
