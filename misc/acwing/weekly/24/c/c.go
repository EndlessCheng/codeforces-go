package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	hs := [4]hp{}
	ms := [4]map[int]*vi{{}, {}, {}, {}}
	var n, m, v int
	Fscan(in, &n)
	p := make([]int, n)
	for i := range p {
		Fscan(in, &p[i])
	}
	for _, p := range p {
		Fscan(in, &v)
		if ms[v][p] == nil {
			ms[v][p] = hs[v].push(p)
		}
	}
	for _, p := range p {
		Fscan(in, &v)
		if ms[v][p] == nil {
			ms[v][p] = hs[v].push(p)
		}
	}

	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &v)
		if len(hs[v]) == 0 {
			Fprint(out, -1, " ")
			continue
		}
		min := hs[v][0].v
		Fprint(out, min, " ")
		for i := 1; i <= 3; i++ {
			if vi := ms[i][min]; vi != nil {
				heap.Remove(&hs[i], vi.i)
				delete(ms[i], min)
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }

type vi struct{ v, i int }
type hp []*vi

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].i = i; h[j].i = j }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(*vi)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v int) *vi     { p := &vi{v, len(*h)}; heap.Push(h, p); return p }
