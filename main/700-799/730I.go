package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf730I(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, p, s, ans int
	Fscan(in, &n, &p, &s)
	oriA := make([]int, n)
	a := make(hp30, n)
	for i := range a {
		Fscan(in, &oriA[i])
		a[i] = pair30{oriA[i], i}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	group := make([]byte, n)
	for _, p := range a[:p] {
		ans += p.v
		group[p.i] = 1
	}
	a = a[p:]
	heap.Init(&a)

	oriB := make([]int, n)
	b := make(hp30, 0, n-p)
	diff := make(hp30, 0, p)
	for i, v := range oriA {
		Fscan(in, &oriB[i])
		if group[i] == 0 {
			b = append(b, pair30{oriB[i], i})
		} else {
			diff = append(diff, pair30{oriB[i] - v, i})
		}
	}
	heap.Init(&b)
	heap.Init(&diff)

	for ; s > 0; s-- {
		for group[a[0].i] > 0 {
			heap.Pop(&a) // 懒删除
		}
		for group[b[0].i] > 0 {
			heap.Pop(&b) // 懒删除
		}
		topA, topB, topD := a[0], b[0], diff[0]
		if topB.v > topA.v+topD.v { // 直接选 b
			ans += topB.v
			group[topB.i] = 2
			heap.Pop(&b)
		} else { // 反悔一个 a 变 b（a 那边选一个更小的）
			ans += topA.v + topD.v
			group[topA.i] = 1
			group[topD.i] = 2
			diff[0] = pair30{oriB[topA.i] - topA.v, topA.i}
			heap.Fix(&diff, 0)
			heap.Pop(&a)
		}
	}

	Fprintln(out, ans)
	for i, g := range group {
		if g == 1 {
			Fprint(out, i+1, " ")
		}
	}
	Fprintln(out)
	for i, g := range group {
		if g == 2 {
			Fprint(out, i+1, " ")
		}
	}
}

//func main() { cf730I(os.Stdin, os.Stdout) }
type pair30 struct{ v, i int }
type hp30 []pair30
func (h hp30) Len() int           { return len(h) }
func (h hp30) Less(i, j int) bool { return h[i].v > h[j].v }
func (h hp30) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp30) Push(v any)        { *h = append(*h, v.(pair30)) }
func (h *hp30) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
