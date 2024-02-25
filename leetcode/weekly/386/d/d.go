package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	if n > m {
		return -1
	}

	total := n
	for _, v := range nums {
		total += v
	}

	firstT := make([]int, n)
	for t := m - 1; t >= 0; t-- {
		firstT[changeIndices[t]-1] = t + 1
	}

	h := hp{}
	ans := n + sort.Search(m+1-n, func(mx int) bool {
		mx += n
		cnt, slow := 0, total
		h = h[:0]
		for t := mx - 1; t >= 0; t-- {
			i := changeIndices[t] - 1
			v := nums[i]
			if v <= 1 || t != firstT[i]-1 {
				cnt++ // 留给左边，用来快速复习/考试
				continue
			}
			if cnt == 0 {
				if len(h) == 0 || v <= h[0].v {
					cnt++ // 留给左边，用来快速复习/考试
					continue
				}
				slow += heap.Pop(&h).(pair).v + 1
				cnt += 2 // 反悔：一天快速复习，一天考试
			}
			slow -= v + 1
			cnt-- // 快速复习，然后消耗一天来考试
			heap.Push(&h, pair{v, i})
		}
		return cnt >= slow // 剩余天数不能快速复习+考试
	})
	if ans > m {
		return -1
	}
	return ans
}

type pair struct{ v, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].v < h[j].v }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
