package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF555B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	var pl, pr, l, r int64
	Fscan(in, &n, &m, &pl, &pr)
	type tuple struct {
		l, r int64
		i    int
	}
	b := make([]tuple, n-1)
	for i := range b {
		Fscan(in, &l, &r)
		b[i] = tuple{l - pr, r - pl, i}
		pl, pr = l, r
	}
	sort.Slice(b, func(i, j int) bool { return b[i].r < b[j].r })
	a := make([]pair55, m)
	for i := range a {
		Fscan(in, &a[i].x)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	ans := make([]int, n-1)
	h, j := hp55{}, 0
	for _, p := range a {
		for ; j < n-1 && b[j].l <= p.x; j++ {
			heap.Push(&h, pair55{b[j].r, b[j].i})
		}
		if len(h) > 0 {
			if h[0].x < p.x {
				break
			}
			ans[heap.Pop(&h).(pair55).i] = p.i
		}
	}
	if len(h) > 0 || j < n-1 {
		Fprint(out, "No")
	} else {
		Fprintln(out, "Yes")
		for _, v := range ans {
			Fprint(out, v, " ")
		}
	}
}

//func main() { CF555B(os.Stdin, os.Stdout) }

type pair55 struct {
	x int64
	i int
}
type hp55 []pair55

func (h hp55) Len() int            { return len(h) }
func (h hp55) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h hp55) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp55) Push(v interface{}) { *h = append(*h, v.(pair55)) }
func (h *hp55) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
