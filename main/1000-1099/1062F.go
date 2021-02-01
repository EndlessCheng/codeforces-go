package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1062F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, ans int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	deg := make([]int, n)
	rg := make([][]int, n)
	rdeg := make([]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		deg[w]++
		rg[w] = append(rg[w], v)
		rdeg[v]++
	}

	cnt := make([]int, n)
	topo := func(g [][]int, deg []int) {
		cntTo := n
		q := []int{}
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
				cntTo--
			}
		}
		for len(q) > 0 {
			v, q = q[0], q[1:]
			if len(q) == 0 {
				cnt[v] += cntTo // 如果队列为空，那么 v 可以到达后面所有的节点
			} else if len(q) == 1 { // 如果队列只剩一个节点，那么判断删去该节点后 v 能否到达后面所有的节点
				for _, w := range g[q[0]] {
					if deg[w] == 1 { // 肥肠抱歉，存在一个可以从 q[0] 达到的点
						goto o
					}
				}
				cnt[v] += cntTo
			}
		o:
			for _, w := range g[v] {
				if deg[w]--; deg[w] == 0 {
					q = append(q, w)
					cntTo--
				}
			}
		}
	}
	topo(g, deg)
	topo(rg, rdeg)
	for _, c := range cnt {
		if c >= n-2 {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1062F(os.Stdin, os.Stdout) }
