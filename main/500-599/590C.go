package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair90 struct{ x, y int }
type deque90 struct{ l, r []pair90 }

func (q deque90) empty() bool     { return len(q.l) == 0 && len(q.r) == 0 }
func (q *deque90) pushL(v pair90) { q.l = append(q.l, v) }
func (q *deque90) pushR(v pair90) { q.r = append(q.r, v) }
func (q *deque90) popL() (v pair90) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func CF590C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	dir4 := []pair90{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, m int
	Fscan(in, &n, &m)
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	bfs := func(tp byte) [][]int {
		dis := make([][]int, n)
		for i := range dis {
			dis[i] = make([]int, m)
			for j := range dis[i] {
				dis[i][j] = 1e8
			}
		}
		q := &deque90{}
		for i, r := range g {
			for j, b := range r {
				if b == tp {
					dis[i][j] = 0
					q.pushL(pair90{i, j})
				}
			}
		}
		for !q.empty() {
			p := q.popL()
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && g[x][y] != '#' {
					d := 0
					if g[x][y] == '.' {
						d = 1
					}
					if newD := dis[p.x][p.y] + d; newD < dis[x][y] {
						dis[x][y] = newD
						if d == 0 {
							q.pushL(pair90{x, y})
						} else {
							q.pushR(pair90{x, y})
						}
					}
				}
			}
		}
		return dis
	}
	d1 := bfs('1')
	d2 := bfs('2')
	d3 := bfs('3')
	ans := int(1e7)
	for i, r := range g {
		for j, b := range r {
			if b != '#' {
				d := d1[i][j] + d2[i][j] + d3[i][j]
				if b == '.' {
					d -= 2
				}
				if d < ans {
					ans = d
				}
			}
		}
	}
	if ans == 1e7 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF590C(os.Stdin, os.Stdout) }
