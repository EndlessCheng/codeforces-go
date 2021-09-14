package main

import "container/heap"

// github.com/EndlessCheng/codeforces-go
type pair struct{ end, left int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

func eatenApples(apples, days []int) (ans int) {
	for i, n, h := 0, len(apples), (hp{}); i < n || len(h) > 0; i++ {
		if i < n {
			h.push(pair{i + days[i], apples[i]})
		}
		for len(h) > 0 && h[0].end <= i {
			h.pop()
		}
		if len(h) > 0 {
			h[0].left--
			if h[0].left == 0 {
				h.pop()
			}
			ans++
		}
	}
	return
}
