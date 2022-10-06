package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF754D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, ans, l int
	Fscan(in, &n, &k)
	a := make([]struct{ l, r, i int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
		a[i].i = i
	}

	sort.Slice(a, func(i, j int) bool { return a[i].l < a[j].l })
	h := hp54{make([]int, k)}
	for i, p := range a[:k] {
		h.IntSlice[i] = p.r
	}
	heap.Init(&h)
	d := h.IntSlice[0] - a[k-1].l + 1
	if d > ans {
		ans = d
		l = a[k-1].l
	}
	for _, p := range a[k:] {
		if p.r > h.IntSlice[0] {
			h.IntSlice[0] = p.r
			heap.Fix(&h, 0)
			d := h.IntSlice[0] - p.l + 1
			if d > ans {
				ans = d
				l = p.l
			}
		}
	}

	Fprintln(out, ans)
	for _, p := range a {
		if ans == 0 || p.l <= l && p.r >= l+ans-1 {
			Fprint(out, p.i+1, " ")
			if k--; k == 0 {
				break
			}
		}
	}
}

//func main() { CF754D(os.Stdin, os.Stdout) }

type hp54 struct{ sort.IntSlice }

func (hp54) Push(interface{})     {}
func (hp54) Pop() (_ interface{}) { return }
