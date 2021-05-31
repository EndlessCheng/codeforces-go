package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func CF1394B(_r io.Reader, out io.Writer) {
	rand.Seed(time.Now().UnixNano())
	in := bufio.NewReader(_r)
	var n, m, k, v, w, wt, total, ans int
	Fscan(in, &n, &m, &k)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		g[v-1] = append(g[v-1], nb{w - 1, wt})
	}

	h := make([]int, n) // int64
	for i := range h {
		h[i] = rand.Int()
		total += h[i]
	}
	sumTo := make([][]int, k+1)
	for i := range sumTo {
		sumTo[i] = make([]int, i)
	}
	for _, es := range g {
		sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })
		for i, e := range es {
			sumTo[len(es)][i] += h[e.to]
		}
	}
	var f func(int, int)
	f = func(p, tot int) {
		if p > k {
			if tot == total {
				ans++
			}
			return
		}
		for _, s := range sumTo[p][:p] {
			f(p+1, tot+s)
		}
	}
	f(1, 0)
	Fprint(out, ans)
}

//func main() { CF1394B(os.Stdin, os.Stdout) }
