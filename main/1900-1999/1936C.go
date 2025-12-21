package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1936C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, ci int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)

		st := n * m
		end := st + n - 1
		type edge struct{ to, wt int }
		g := make([][]edge, end+1)

		for i := range n {
			Fscan(in, &ci)
			v := n*m + i
			for w := i * m; w < (i+1)*m; w++ {
				g[v] = append(g[v], edge{w, 0})
				g[w] = append(g[w], edge{v, ci})
			}
		}

		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}
		type pair struct{ v, i int }
		p := make([]pair, n)
		for j := range m {
			for i, r := range a {
				p[i] = pair{r[j], i}
			}
			slices.SortFunc(p, func(a, b pair) int { return a.v - b.v })
			for i := range n - 1 {
				v := p[i].i*m + j
				w := p[i+1].i*m + j
				g[v] = append(g[v], edge{w, 0})
				g[w] = append(g[w], edge{v, p[i+1].v - p[i].v})
			}
		}

		dis := make([]int, len(g))
		for i := range dis {
			dis[i] = 1e18
		}
		dis[st] = 0
		h := hp36{{0, st}}
		for len(h) > 0 {
			p := heap.Pop(&h).(vd36)
			if p.d > dis[p.v] {
				continue
			}
			for _, e := range g[p.v] {
				w := e.to
				d := p.d + e.wt
				if d < dis[w] {
					dis[w] = d
					heap.Push(&h, vd36{d, w})
				}
			}
		}
		Fprintln(out, dis[end])
	}
}

//func main() { cf1936C(bufio.NewReader(os.Stdin), os.Stdout) }

type vd36 struct{ d, v int }
type hp36 []vd36

func (h hp36) Len() int           { return len(h) }
func (h hp36) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp36) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp36) Push(v any)        { *h = append(*h, v.(vd36)) }
func (h *hp36) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
