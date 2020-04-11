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
	type pair struct{ x, y int }
	dir4 := [...]pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

	solve := func(_case int) (ans int) {
		var n, m int
		Fscan(in, &n, &m)
		q := []pair{}            // 检查队列
		onQ := make([][]bool, n) // 记录人 (x,y) 是否在下一轮的检查队列中
		mat := make([][]int, n)
		nei := make([][][4]pair, n)
		curSum := 0
		for i := range mat {
			onQ[i] = make([]bool, m)
			mat[i] = make([]int, m)
			nei[i] = make([][4]pair, m)
			for j := range mat[i] {
				Fscan(in, &mat[i][j])
				curSum += mat[i][j]
				for k, d := range dir4 {
					nei[i][j][k] = pair{i + d.x, j + d.y}
				}
				q = append(q, pair{i, j})
			}
		}

		for {
			ans += curSum
			removes := []pair{} // 要被淘汰的人
			curQ := q
			q = []pair{}
			for _, p := range curQ {
				i, j := p.x, p.y
				v := mat[i][j]
				if v == 0 { // 要检查的人已经被淘汰了
					continue
				}
				s, cnt := 0, 0
				for _, p := range nei[i][j] {
					if x, y := p.x, p.y; x >= 0 && x < n && y >= 0 && y < m {
						s += mat[x][y]
						cnt++
					}
				}
				if cnt*v < s {
					removes = append(removes, pair{i, j}) // 淘汰
					for _, p := range nei[i][j] {
						if x, y := p.x, p.y; x >= 0 && x < n && y >= 0 && y < m && !onQ[x][y] {
							q = append(q, pair{x, y}) // 下一轮要检查的人
							onQ[x][y] = true
						}
					}
				}
			}

			if len(removes) == 0 {
				break
			}

			for _, p := range removes {
				curSum -= mat[p.x][p.y]
				mat[p.x][p.y] = 0
			}

			for _, p := range q {
				i, j := p.x, p.y
				onQ[i][j] = false
				for k := range nei[i][j] {
					// 不断循环去找一个未被淘汰的人
					for {
						p := nei[i][j][k]
						if x, y := p.x, p.y; x >= 0 && x < n && y >= 0 && y < m && mat[x][y] == 0 {
							nei[i][j][k] = nei[x][y][k]
						} else {
							break
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
		Fprintf(out, "Case #%d: %d\n", _case, solve(_case))
	}
}

func main() { run(os.Stdin, os.Stdout) }
