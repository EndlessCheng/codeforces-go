package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type deque42 struct{ l, r []int }

func (q deque42) empty() bool  { return len(q.l) == 0 && len(q.r) == 0 }
func (q *deque42) pushL(v int) { q.l = append(q.l, v) }
func (q *deque42) pushR(v int) { q.r = append(q.r, v) }
func (q *deque42) popL() (v int) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func CF1442C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	pow := func(x int64, n int) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var n, m, v, w int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v] = append(g[v], nb{w, 0})
		g[w] = append(g[w], nb{v, 1})
	}

	minTr := make([]int, n+1)
	for i := range minTr {
		minTr[i] = 1e9
	}
	minTr[1] = 0
	q := &deque42{}
	q.pushL(1)
	for !q.empty() {
		v := q.popL()
		for _, e := range g[v] {
			w, d := e.to, minTr[v]&1^e.wt // 根据 minTr[v] 决定实际方向
			if newD := minTr[v] + d; newD < minTr[w] {
				minTr[w] = newD
				if d == 0 {
					q.pushL(w)
				} else {
					q.pushR(w)
				}
			}
		}
	}

	dis := make([][20]int, n+1)
	for i := range dis {
		for j := range dis[i] {
			dis[i][j] = 1e9
		}
	}
	dis[1][0] = 0
	type pair struct{ v, lv int }
	q2 := []pair{{1, 0}}
	for len(q2) > 0 {
		p := q2[0]
		q2 = q2[1:]
		v := p.v
		for _, e := range g[v] {
			w, wt := e.to, p.lv&1^e.wt // 根据 p.lv 决定实际方向
			// 通过减去 minTr 来得到映射到 dis 上的层数
			if lv := p.lv + wt - minTr[w]; lv < 20 {
				if newD := dis[v][p.lv-minTr[v]] + 1; newD < dis[w][lv] {
					dis[w][lv] = newD
					q2 = append(q2, pair{w, p.lv + wt})
				}
			}
		}
	}

	minT, minD := int(1e9), 0
	for i, d := range dis[n] {
		// t 为实际层数
		if t := i + minTr[n]; minT > 19 || t > 19 {
			if t < minT {
				minT, minD = t, d
			}
		} else if 1<<t+d < 1<<minT+minD {
			minT, minD = t, d
		}
	}
	Fprint(out, (int64(minD)+pow(2, minT)-1)%mod)
}

//func main() { CF1442C(os.Stdin, os.Stdout) }
