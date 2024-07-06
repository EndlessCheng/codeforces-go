package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"math"
)

func cf1725M(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	type edge struct{ to, wt, inv int }
	g := make([][]edge, n)
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], edge{w, wt, 0})
		g[w] = append(g[w], edge{v, wt, 1})
	}

	dis := make([][2]int, n)
	for i := 1; i < n; i++ {
		dis[i] = [2]int{math.MaxInt, math.MaxInt}
	}
	h := hp25{{}, {inv: 1}}
	for len(h) > 0 {
		p := heap.Pop(&h).(data25)
		if p.dis > dis[p.v][p.inv] {
			continue
		}
		for _, e := range g[p.v] {
			w, eInv := e.to, e.inv
			newD := p.dis + e.wt
			if (eInv == 1 || p.inv == 0) && newD < dis[w][eInv] {
				dis[w][eInv] = newD
				heap.Push(&h, data25{newD, w, eInv})
			}
		}
	}
	for _, d := range dis[1:] {
		res := min(d[0], d[1])
		if res == math.MaxInt {
			res = -1
		}
		Fprint(out, res, " ")
	}
}

//func main() { cf1725M(bufio.NewReader(os.Stdin), os.Stdout) }
type data25 struct{ dis, v, inv int }
type hp25 []data25
func (h hp25) Len() int           { return len(h) }
func (h hp25) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp25) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp25) Push(v any)        { *h = append(*h, v.(data25)) }
func (h *hp25) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
