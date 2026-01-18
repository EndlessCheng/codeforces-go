package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxCapacity1(costs []int, capacity []int, budget int) (ans int) {
	// 把 costs[i] 和 capacity[i] 绑在一起排序
	type pair struct{ cost, cap int }
	n := len(costs)
	a := make([]pair, 0, n)
	for i, cost := range costs {
		if cost < budget { // 太贵的机器直接忽略
			a = append(a, pair{cost, capacity[i]})
		}
	}
	slices.SortFunc(a, func(a, b pair) int { return a.cost - b.cost })

	preMax := make([]int, len(a)+1)
	for i, p := range a {
		// 二分第一台价格 >= budget-p.cost 的机器，下标减一，就是最后一台价格 < budget-p.cost 的机器
		j := sort.Search(i, func(j int) bool { return a[j].cost >= budget-p.cost })
		// (j - 1) + 1 == j
		ans = max(ans, p.cap+preMax[j]) // j=0 的情况对应单选一台机器
		preMax[i+1] = max(preMax[i], p.cap)
	}
	return
}

func maxCapacity(costs, capacity []int, budget int) (ans int) {
	type pair struct{ cost, cap int }
	n := len(costs)
	a := make([]pair, 0, n)
	for i, cost := range costs {
		if cost < budget {
			a = append(a, pair{cost, capacity[i]})
		}
	}
	slices.SortFunc(a, func(a, b pair) int { return a.cost - b.cost })

	st := []pair{{}} // 栈底加个哨兵
	for _, p := range a {
		j := sort.Search(len(st), func(j int) bool { return st[j].cost >= budget-p.cost }) - 1
		ans = max(ans, p.cap+st[j].cap) // j=0 的情况对应单选一台机器
		if p.cap > st[len(st)-1].cap {
			st = append(st, p)
		}
	}
	return
}
