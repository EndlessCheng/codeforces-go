package main

import "github.com/emirpasic/gods/v2/trees/redblacktree"

func avoidFlood1(rains []int) []int {
	ans := make([]int, len(rains))
	fullDay := map[int]int{}                    // lake -> 装满日
	dryDay := redblacktree.New[int, struct{}]() // 未被使用的抽水日
	for i, lake := range rains {
		if lake == 0 {
			ans[i] = 1                // 先随便选一个湖抽干
			dryDay.Put(i, struct{}{}) // 保存抽水日
			continue
		}
		if j, ok := fullDay[lake]; ok {
			// 必须在 j 之后，i 之前把 lake 抽干
			// 选一个最早的未被使用的抽水日，如果选晚的，可能会导致其他湖没有可用的抽水日
			node, _ := dryDay.Ceiling(j)
			if node == nil {
				return nil // 无法阻止洪水
			}
			ans[node.Key] = lake
			dryDay.Remove(node.Key) // node.Key 已使用，移除
		}
		ans[i] = -1
		fullDay[lake] = i // 插入或更新装满日
	}
	return ans
}

func avoidFlood(rains []int) []int {
	n := len(rains)
	// 非递归并查集
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	ans := make([]int, n)
	fullDay := map[int]int{} // lake -> 装满日
	for i, lake := range rains {
		if lake == 0 {
			ans[i] = 1 // 先随便选一个湖抽干
			continue
		}
		if j, ok := fullDay[lake]; ok {
			// 必须在 j 之后，i 之前把 lake 抽干
			// 选一个最早的未被使用的抽水日，如果选晚的，可能会导致其他湖没有可用的抽水日
			dryDay := find(j + 1)
			if dryDay >= i {
				return nil // 无法阻止洪水
			}
			ans[dryDay] = lake
			fa[dryDay] = find(dryDay + 1) // 删除 dryDay
		}
		ans[i] = -1
		fa[i] = i + 1 // 删除 i
		fullDay[lake] = i // 插入或更新装满日
	}
	return ans
}
