package copypasta

import "sort"

// 最小堆
type intHeap struct {
	sort.IntSlice
}

func (h *intHeap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *intHeap) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

//

type hPair struct {
	x int64
	y int
}
type pairHeap []hPair

func (h pairHeap) Len() int              { return len(h) }
func (h pairHeap) Less(i, j int) bool    { return h[i].x < h[j].x || h[i].x == h[j].x && h[i].y < h[j].y }
func (h pairHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(v interface{})   { *h = append(*h, v.(hPair)) }
func (h *pairHeap) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

//

type prioQueue struct {
	heap  []int // vertices in heap order
	index []int // index of each vertex in the heap
	cost  []int
}

func emptyPrioQueue(cost []int) *prioQueue {
	return &prioQueue{
		index: make([]int, len(cost)),
		cost:  cost,
	}
}

func newPrioQueue(cost []int) *prioQueue {
	n := len(cost)
	q := &prioQueue{
		heap:  make([]int, n),
		index: make([]int, n),
		cost:  cost,
	}
	for i := range q.heap {
		q.heap[i] = i
		q.index[i] = i
	}
	return q
}

// Len returns the number of elements in the queue.
func (q *prioQueue) Len() int {
	return len(q.heap)
}

// Push pushes v onto the queue.
// The time complexity is O(log n) where n = q.Len().
func (q *prioQueue) Push(v int) {
	n := q.Len()
	q.heap = append(q.heap, v)
	q.index[v] = n
	q.up(n)
}

// Pop removes the minimum element from the queue and returns it.
// The time complexity is O(log n) where n = q.Len().
func (q *prioQueue) Pop() int {
	n := q.Len() - 1
	q.swap(0, n)
	q.down(0, n)

	v := q.heap[n]
	q.index[v] = -1
	q.heap = q.heap[:n]
	return v
}

// Contains tells whether v is in the queue.
func (q *prioQueue) Contains(v int) bool {
	return q.index[v] >= 0
}

// Fix re-establishes the ordering after the cost for v has changed.
// The time complexity is O(log n) where n = q.Len().
func (q *prioQueue) Fix(v int) {
	if i := q.index[v]; !q.down(i, q.Len()) {
		q.up(i)
	}
}

func (q *prioQueue) less(i, j int) bool {
	return q.cost[q.heap[i]] < q.cost[q.heap[j]]
}

func (q *prioQueue) swap(i, j int) {
	q.heap[i], q.heap[j] = q.heap[j], q.heap[i]
	q.index[q.heap[i]] = i
	q.index[q.heap[j]] = j
}

func (q *prioQueue) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !q.less(j, i) {
			break
		}
		q.swap(i, j)
		j = i
	}
}

func (q *prioQueue) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && q.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !q.less(j, i) {
			break
		}
		q.swap(i, j)
		i = j
	}
	return i > i0
}
