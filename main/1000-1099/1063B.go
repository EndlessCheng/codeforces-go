package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair63 struct{ x, y, l, r int }
type deque63 struct{ l, r []pair63 }

func (q deque63) empty() bool     { return len(q.l) == 0 && len(q.r) == 0 }
func (q *deque63) pushL(v pair63) { q.l = append(q.l, v) }
func (q *deque63) pushR(v pair63) { q.r = append(q.r, v) }
func (q *deque63) popL() (v pair63) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func CF1063B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	dir4 := []struct{ x, y, l, r int }{{x: -1}, {x: 1}, {0, -1, 1, 0}, {0, 1, 0, 1}}

	var n, m, sx, sy, ll, rr, ans int
	Fscan(in, &n, &m, &sx, &sy, &ll, &rr)
	sx--
	sy--
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}

	q := deque63{}
	q.pushL(pair63{sx, sy, 0, 0})
	g[sx][sy] = 0
	for !q.empty() {
		ans++
		p := q.popL()
		for _, d := range dir4 {
			x, y, l, r := p.x+d.x, p.y+d.y, p.l+d.l, p.r+d.r
			if 0 <= x && x < n && 0 <= y && y < m && l <= ll && r <= rr && g[x][y] == '.' {
				g[x][y] = 0
				p := pair63{x, y, l, r}
				if d.y == 0 {
					q.pushL(p)
				} else {
					q.pushR(p)
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1063B(os.Stdin, os.Stdout) }
