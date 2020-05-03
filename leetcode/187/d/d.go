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

	sum0 := 0
	for _, row := range mat {
		sum0 += row[0]
	}
	q.push(pair{sum: sum0})
	vis[[40]int8{}] = true
	for {
		p := q.pop()
		sum, pos := p.sum, p.pos
		k--
		if k == 0 {
			return sum
		}
		for i, row := range mat {
			if pos[i]+1 == m {
				continue
			}
			next := pos
			next[i]++
			if !vis[next] {
				vis[next] = true
				q.push(pair{sum + row[next[i]] - row[next[i]-1], next})
			}
		}
	}
}
