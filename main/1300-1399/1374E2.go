package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1374E2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m, k, t, t0, t1, s int
	g := [4][]pair74{}
	Fscan(in, &n, &m, &k)
	for i := 1; i <= n; i++ {
		Fscan(in, &t, &t0, &t1)
		x := t0<<1 | t1
		g[x] = append(g[x], pair74{t, i})
	}
	trash, a, b, both := g[0], g[1], g[2], g[3]
	if len(a) > len(b) {
		a, b = b, a
	}
	na, nb := len(a), len(both)
	if na+nb < k || nb < k*2-m {
		Fprint(out, -1)
		return
	}

	Sort := func(a []pair74) { sort.Slice(a, func(i, j int) bool { return a[i].t < a[j].t }) }
	Sort(a)
	Sort(b)
	limitNa := min(m/2, m-k)
	if na > limitNa {
		trash = append(append(trash, a[limitNa:]...), b[limitNa:]...)
		a = a[:limitNa]
		b = b[:limitNa]
		na = limitNa
	} else {
		trash = append(trash, b[na:]...)
		b = b[:na]
	}
	for i, p := range a {
		s += p.t + b[i].t
	}

	Sort(both)
	b0 := max(max(k-na, 0), m-2*na-len(trash))
	for _, p := range both[:b0] {
		s += p.t
	}

	Sort(trash)
	h := maxMinHeap{&maxHp{}, &minHp{}}
	for _, p := range trash[:m-2*na-b0] {
		h.left.push(p)
	}
	for _, p := range trash[m-2*na-b0:] {
		h.right.push(p)
	}

	s0 := s
	tmp := minHp(append([]pair74{}, *h.right...))
	h0 := maxMinHeap{&maxHp{append(h.left.a[:0:0], h.left.a...), h.left.s}, &tmp}

	ans := s + h.left.s
	for i := na - 1; i >= k; i-- {
		s -= a[i].t + b[i].t
		h.push(a[i])
		h.push(b[i])
		ans = min(ans, s+h.left.s)
	}

	if len(both) > m {
		both = both[:m]
	}
	for j, p := range both[b0:] {
		s += p.t
		i := min(na, k) - 1 - j
		if i >= 0 {
			s -= a[i].t + b[i].t
			h.push(a[i])
			h.push(b[i])
		}
		h.l2r()
		ans = min(ans, s+h.left.s)
	}
	Fprintln(out, ans)

	output := func() {
		for _, p := range a {
			Fprint(out, p.i, " ")
		}
		for _, p := range b {
			Fprint(out, p.i, " ")
		}
		for _, p := range both {
			Fprint(out, p.i, " ")
		}
		for _, p := range h.left.a {
			Fprint(out, p.i, " ")
		}
	}

	s = s0
	h = h0

	if s+h.left.s == ans {
		both = both[:b0]
		output()
		return
	}

	for i := na - 1; i >= k; i-- {
		s -= a[i].t + b[i].t
		h.push(a[i])
		h.push(b[i])
		if s+h.left.s == ans {
			a = a[:i]
			b = b[:i]
			both = both[:b0]
			output()
			return
		}
	}

	for j, p := range both[b0:] {
		s += p.t
		i := min(na, k) - 1 - j
		if i >= 0 {
			s -= a[i].t + b[i].t
			h.push(a[i])
			h.push(b[i])
		}
		h.l2r()
		if s+h.left.s == ans {
			if i > 0 {
				a = a[:i]
				b = b[:i]
			} else {
				a = nil
				b = nil
			}
			both = both[:b0+j+1]
			output()
			return
		}
	}
}

//func main() { CF1374E2(os.Stdin, os.Stdout) }

type maxMinHeap struct {
	left  *maxHp
	right *minHp
}

func (h maxMinHeap) push(v pair74) { h.left.push(h.right.pushPop(v)) }
func (h maxMinHeap) l2r()          { h.right.push(heap.Pop(h.left).(pair74)) }
func (h maxMinHeap) r2l()          { h.left.push(heap.Pop(h.right).(pair74)) }

type pair74 struct{ t, i int }
type maxHp struct {
	a []pair74
	s int
}

func (h maxHp) Len() int            { return len(h.a) }
func (h maxHp) Less(i, j int) bool  { return h.a[i].t > h.a[j].t }
func (h maxHp) Swap(i, j int)       { h.a[i], h.a[j] = h.a[j], h.a[i] }
func (h *maxHp) Push(v interface{}) { h.s += v.(pair74).t; h.a = append(h.a, v.(pair74)) }
func (h *maxHp) Pop() interface{}   { v := h.a[len(h.a)-1]; h.s -= v.t; h.a = h.a[:len(h.a)-1]; return v }
func (h *maxHp) push(v pair74)      { heap.Push(h, v) }
func (h *maxHp) pushPop(v pair74) pair74 {
	if h.Len() > 0 && v.t < h.a[0].t {
		h.s += v.t - h.a[0].t
		v, h.a[0] = h.a[0], v
		heap.Fix(h, 0)
	}
	return v
}

type minHp []pair74

func (h minHp) Len() int            { return len(h) }
func (h minHp) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h minHp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHp) Push(v interface{}) { *h = append(*h, v.(pair74)) }
func (h *minHp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *minHp) push(v pair74)      { heap.Push(h, v) }
func (h minHp) pushPop(v pair74) pair74 {
	if h.Len() > 0 && v.t > h[0].t {
		v, h[0] = h[0], v
		heap.Fix(&h, 0)
	}
	return v
}
