package main

import "container/heap"

func minimumPairRemoval(nums []int) (ans int) {
	n := len(nums)
	h := hp{}
	dec := 0 // 递减的相邻对的个数
	for i := range n - 1 {
		x, y := nums[i], nums[i+1]
		if x > y {
			dec++
		}
		h = append(h, pair{x + y, i})
	}
	heap.Init(&h)
	lazy := map[pair]int{}

	// 每个下标的左右最近的未删除下标
	left := make([]int, n+1) // 加一个哨兵，防止下标越界
	right := make([]int, n)
	for i := range n {
		left[i] = i - 1
		right[i] = i + 1
	}
	remove := func(i int) {
		l, r := left[i], right[i]
		right[l] = r
		left[r] = l
	}

	for dec > 0 {
		ans++

		for lazy[h[0]] > 0 {
			lazy[h[0]]--
			heap.Pop(&h)
		}
		p := heap.Pop(&h).(pair) // 删除相邻元素和最小的一对
		s := p.s
		i := p.i

		// (当前元素，下一个数)
		nxt := right[i]
		if nums[i] > nums[nxt] { // 旧数据
			dec--
		}

		// (前一个数，当前元素)
		pre := left[i]
		if pre >= 0 {
			if nums[pre] > nums[i] { // 旧数据
				dec--
			}
			if nums[pre] > s { // 新数据
				dec++
			}
			lazy[pair{nums[pre] + nums[i], pre}]++ // 懒删除
			heap.Push(&h, pair{nums[pre] + s, pre})
		}

		// (下一个数，下下一个数)
		nxt2 := right[nxt]
		if nxt2 < n {
			if nums[nxt] > nums[nxt2] { // 旧数据
				dec--
			}
			if s > nums[nxt2] { // 新数据（当前元素，下下一个数）
				dec++
			}
			lazy[pair{nums[nxt] + nums[nxt2], nxt}]++ // 懒删除
			heap.Push(&h, pair{s + nums[nxt2], i})
		}

		nums[i] = s
		remove(nxt)
	}
	return
}

type pair struct{ s, i int } // (相邻元素和，左边那个数的下标)
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.s < b.s || a.s == b.s && a.i < b.i }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
