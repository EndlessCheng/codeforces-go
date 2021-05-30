package main

import "container/heap"

// github.com/EndlessCheng/codeforces-go
type pr struct{ v, i int }
type hp []pr

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.v < b.v || a.v == b.v && a.i < b.i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pr)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pr)          { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(pr).i }

func assignTasks(servers, tasks []int) []int {
	ans := make([]int, len(tasks))
	idle := make(hp, len(servers))
	for i, s := range servers {
		idle[i] = pr{s, i}
	}
	heap.Init(&idle)
	busy := hp{}
	cur := 0
	release := func() {
		for len(busy) > 0 && busy[0].v <= cur {
			id := busy.pop()
			idle.push(pr{servers[id], id})
		}
	}
	for i, t := range tasks {
		if i > cur {
			cur = i
		}
		release()
		if len(idle) == 0 {
			cur = busy[0].v
			release()
		}
		id := idle.pop()
		ans[i] = id
		busy.push(pr{cur + t, id})
	}
	return ans
}
