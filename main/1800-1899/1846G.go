package main

import (
	"container/heap"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func bin46(s string) (a int) {
	for i, b := range s {
		a |= int(b&1) << i
	}
	return
}

func cf1846G(in io.Reader, out io.Writer) {
	var T, n, m int
	var s0, s, t string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s0)
		a := make([]struct{ d, s, t int }, m)
		for i := range a {
			Fscan(in, &a[i].d, &s, &t)
			a[i].s = bin46(s)
			a[i].t = bin46(t)
		}

		dis := make([]int, 1<<n)
		for i := range dis {
			dis[i] = 1e9
		}
		st := bin46(s0)
		dis[st] = 0
		h := hp46{{0, st}}
		for len(h) > 0 {
			p := heap.Pop(&h).(pair46)
			d := p.d
			v := p.v
			if v == 0 {
				Fprintln(out, d)
				continue o
			}
			if d > dis[v] {
				continue
			}
			for _, t := range a {
				newD := d + t.d
				w := v&^t.s | t.t
				if newD < dis[w] {
					dis[w] = newD
					heap.Push(&h, pair46{newD, w})
				}
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { cf1846G(bufio.NewReader(os.Stdin), os.Stdout) }

type pair46 struct{ d, v int }
type hp46 []pair46
func (h hp46) Len() int           { return len(h) }
func (h hp46) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp46) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp46) Push(v any)        { *h = append(*h, v.(pair46)) }
func (h *hp46) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
