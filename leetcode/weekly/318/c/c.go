package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func totalCost(costs []int, k, candidates int) int64 {
	ans := 0
	if n := len(costs); candidates*2 < n {
		pre := hp{costs[:candidates]}
		heap.Init(&pre)
		suf := hp{costs[n-candidates:]}
		heap.Init(&suf)
		for i, j := candidates, n-1-candidates; k > 0 && i <= j; k-- {
			if pre.IntSlice[0] <= suf.IntSlice[0] {
				ans += pre.IntSlice[0]
				pre.IntSlice[0] = costs[i]
				i++
				heap.Fix(&pre, 0)
			} else {
				ans += suf.IntSlice[0]
				suf.IntSlice[0] = costs[j]
				j--
				heap.Fix(&suf, 0)
			}
		}
		costs = append(pre.IntSlice, suf.IntSlice...)
	}
	sort.Ints(costs)
	for _, v := range costs[:k] { // 也可以用快速选择算法求前 k 小
		ans += v
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
