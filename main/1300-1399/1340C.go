package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type pair40 struct{ i, t int }
type deque40 struct{ l, r []pair40 }

func (q deque40) empty() bool     { return len(q.l) == 0 && len(q.r) == 0 }
func (q *deque40) pushL(v pair40) { q.l = append(q.l, v) }
func (q *deque40) pushR(v pair40) { q.r = append(q.r, v) }
func (q *deque40) popL() (v pair40) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func CF1340C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var n, m, g, r int
	Fscan(in, &n, &m)
	x := make([]int, m)
	for i := range x {
		Fscan(in, &x[i])
	}
	sort.Ints(x)
	Fscan(in, &g, &r)
	for i := 1; i < m; i++ {
		if x[i]-x[i-1] > g {
			Fprint(out, -1)
			return
		}
	}

	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, g+1)
		for j := range dis[i] {
			dis[i][j] = 1e9
		}
	}
	dis[0][g] = 0
	q := &deque40{}
	q.pushL(pair40{0, g})
	for !q.empty() {
		p := q.popL()
		i, t := p.i, p.t
		d := dis[i][t]
		if t > 0 {
			if i > 0 {
				if j := t - (x[i] - x[i-1]); j >= 0 && d < dis[i-1][j] {
					dis[i-1][j] = d
					q.pushL(pair40{i - 1, j})
				}
			}
			if i < m-1 {
				if j := t - (x[i+1] - x[i]); j >= 0 && d < dis[i+1][j] {
					dis[i+1][j] = d
					q.pushL(pair40{i + 1, j})
				}
			}
		} else {
			newD := d + 1
			if i > 0 {
				if j := g - (x[i] - x[i-1]); newD < dis[i-1][j] {
					dis[i-1][j] = newD
					q.pushR(pair40{i - 1, j})
				}
			}
			if i < m-1 {
				if j := g - (x[i+1] - x[i]); newD < dis[i+1][j] {
					dis[i+1][j] = newD
					q.pushR(pair40{i + 1, j})
				}
			}
		}
	}
	ans := int64(1e18)
	for j, d := range dis[m-1] {
		if d < 1e9 {
			ans = min(ans, int64(d)*int64(r+g)+int64(g-j))
		}
	}
	if ans == 1e18 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF1340C(os.Stdin, os.Stdout) }
