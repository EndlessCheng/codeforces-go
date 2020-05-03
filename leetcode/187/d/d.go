package main

import "container/heap"

type pair struct {
	sum int
	pos [40]int8
}
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].sum < h[j].sum }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v pair)          { heap.Push(h, v) }
func (h *hp) pop() pair            { return heap.Pop(h).(pair) }
func (h hp) empty() bool           { return len(h) == 0 }

func kthSmallest(mat [][]int, k int) (ans int) {
	m := int8(len(mat[0]))
	q := &hp{}
	vis := map[[40]int8]bool{}

	sum := 0
	for _, row := range mat {
		sum += row[0]
	}
	q.push(pair{sum: sum})
	vis[[40]int8{}] = true
	for !q.empty() {
		p := q.pop()
		k--
		if k == 0 {
			return p.sum
		}
		for i, row := range mat {
			if p.pos[i]+1 == m {
				continue
			}
			next := p.pos
			next[i]++
			if !vis[next] {
				vis[next] = true
				q.push(pair{p.sum + row[next[i]] - row[next[i]-1], next})
			}
		}
	}
	panic(-1)
}
