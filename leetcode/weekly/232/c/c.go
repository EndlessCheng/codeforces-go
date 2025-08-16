package main

import "container/heap"

// github.com/EndlessCheng/codeforces-go
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	n := len(classes)
	h := make(hp, n)
	for i, c := range classes {
		h[i] = pair{c[0], c[1]}
	}
	heap.Init(&h)

	for range extraStudents {
		h[0].a++
		h[0].b++
		heap.Fix(&h, 0)
	}

	sum := 0.0
	for _, t := range h {
		sum += float64(t.a) / float64(t.b)
	}
	return sum / float64(n)
}

type pair struct{ a, b int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return (a.b-a.a)*b.b*(b.b+1) > (b.b-b.a)*a.b*(a.b+1)
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)        {}
func (hp) Pop() (_ any)    { return }
