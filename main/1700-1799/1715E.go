package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"math/big"
)

// https://github.com/EndlessCheng
type vec15 struct{ x, y int }

func (a vec15) sub(b vec15) vec15 { return vec15{a.x - b.x, a.y - b.y} }
func (a vec15) dot(b vec15) int   { return a.x*b.x + a.y*b.y }
func (a vec15) det(b vec15) bool {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w) <= 0
}

func cf1715E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, k int
	Fscan(in, &n, &m, &k)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for range m {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = 1e18
	}
	dis[0] = 0
	dij := func() []int {
		h := make(hp15, n)
		for i, d := range dis {
			h[i] = pair15{d, i}
		}
		heap.Init(&h)
		for len(h) > 0 {
			p := heap.Pop(&h).(pair15)
			v := p.v
			d := p.dis
			if d > dis[v] {
				continue
			}
			for _, e := range g[v] {
				w := e.to
				newD := d + e.wt
				if newD < dis[w] {
					dis[w] = newD
					heap.Push(&h, pair15{newD, w})
				}
			}
		}
		return dis
	}

	for range k {
		dij()
		q := []vec15{}
		for i, d := range dis {
			v := vec15{i, i*i + d}
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(v.sub(q[len(q)-1])) {
				q = q[:len(q)-1]
			}
			q = append(q, v)
		}
		for i, d := range dis {
			p := vec15{-2 * i, 1}
			for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
				q = q[1:]
			}
			dis[i] = min(d, p.dot(q[0])+i*i)
		}
	}
	dij()

	for _, v := range dis {
		Fprint(out, v, " ")
	}
}

//func main() { cf1715E(bufio.NewReader(os.Stdin), os.Stdout) }

type pair15 struct{ dis, v int }
type hp15 []pair15
func (h hp15) Len() int           { return len(h) }
func (h hp15) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp15) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp15) Push(v any)        { *h = append(*h, v.(pair15)) }
func (h *hp15) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
