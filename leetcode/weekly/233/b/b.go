package main

import "container/heap"

// github.com/EndlessCheng/codeforces-go
type pair struct{ price, left int }
type hp []pair
type hp2 []pair

func (h hp) Len() int             { return len(h) }
func (h hp) Less(i, j int) bool   { return h[i].price > h[j].price }
func (h hp) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})  { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}    { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)         { heap.Push(h, v) }
func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp2) push(v pair)        { heap.Push(h, v) }

func getNumberOfBacklogOrders(orders [][]int) (ans int) {
	h0, h1 := hp{}, hp2{}
	for _, o := range orders {
		price, left := o[0], o[1]
		if o[2] == 0 {
			for left > 0 && len(h1) > 0 && h1[0].price <= price {
				if h1[0].left > left {
					h1[0].left -= left
					left = 0
					break
				}
				left -= h1[0].left
				heap.Pop(&h1)
			}
			if left > 0 {
				h0.push(pair{price, left})
			}
		} else {
			for left > 0 && len(h0) > 0 && h0[0].price >= price {
				if h0[0].left > left {
					h0[0].left -= left
					left = 0
					break
				}
				left -= h0[0].left
				heap.Pop(&h0)
			}
			if left > 0 {
				h1.push(pair{price, left})
			}
		}
	}
	for _, p := range h0 {
		ans += p.left
	}
	for _, p := range h1 {
		ans += p.left
	}
	return ans % (1e9 + 7)
}
