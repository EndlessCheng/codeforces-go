package main

import (
	"bufio"
	. "fmt"
	"os"
)

type (
	input1174 struct {
		n int
		g [][]int
	}
	guess1174 struct{ ans int }
	qIn1174   struct {
		tp string
		v  int
	}
	qOut1174 struct{ v int }
)

// github.com/EndlessCheng/codeforces-go
func CF1174F(in input1174, Q func(qIn1174) qOut1174) (gs guess1174) {
	n, g := in.n, in.g
	type node struct{ d, sz, hson int }
	nodes := make([]node, n+1)
	var f func(v, fa, d int) *node
	f = func(v, fa, d int) *node {
		nodes[v] = node{d, 1, -1}
		o := &nodes[v]
		for _, w := range g[v] {
			if w != fa {
				son := f(w, v, d+1)
				o.sz += son.sz
				if o.hson == -1 || son.sz > nodes[o.hson].sz {
					o.hson = w
				}
			}
		}
		return o
	}
	f(1, 0, 0)

	depX := Q(qIn1174{"d", 1}).v
	var f2 func(v int) int
	f2 = func(v int) int {
		hp := []int{v}
		for o := nodes[v]; o.hson != -1; o = nodes[o.hson] {
			hp = append(hp, o.hson)
		}
		end := hp[len(hp)-1]
		depAX := (depX + nodes[end].d - Q(qIn1174{"d", end}).v) / 2
		ax := hp[depAX-nodes[v].d]
		if depAX == depX {
			return ax
		}
		return f2(Q(qIn1174{"s", ax}).v)
	}
	gs.ans = f2(1)
	return
}

func ioq1174() {
	in := bufio.NewReader(os.Stdin)
	Q := func(q qIn1174) (resp qOut1174) { Println(q.tp, q.v); Fscan(in, &resp.v); return }
	d := input1174{}
	Fscan(in, &d.n)
	g := make([][]int, d.n+1)
	for i := 1; i < d.n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	d.g = g
	gs := CF1174F(d, Q)
	Println("!", gs.ans)
}

//func main() { ioq() }
