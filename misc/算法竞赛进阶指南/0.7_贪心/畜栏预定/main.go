package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P2859

// github.com/EndlessCheng/codeforces-go
type pair struct{ r, i int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].r < h[j].r }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (_ interface{}) { return }
func (h *hp) push(r int)           { heap.Push(h, pair{r, h.Len()}) }

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]struct{ l, r, i int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })

	ans := make([]int, n)
	h := hp{}
	for _, p := range a {
		if len(h) == 0 || p.l <= h[0].r {
			ans[p.i] = len(h)
			h.push(p.r)
		} else {
			ans[p.i] = h[0].i
			h[0].r = p.r
			heap.Fix(&h, 0)
		}
	}
	Fprintln(out, h.Len())
	for _, v := range ans {
		Fprintln(out, v+1)
	}
}

func main() { run(os.Stdin, os.Stdout) }
