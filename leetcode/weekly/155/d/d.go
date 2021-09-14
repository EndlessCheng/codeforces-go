package main

func topSort(g [][]int, deg, items []int) (orders []int) {
	q := []int{}
	for _, i := range items {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		orders = append(orders, v)
		for _, w := range g[v] {
			deg[w]--
			if deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	return
}

func sortItems(n, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupG := make([][]int, m+n)
	groupDeg := make([]int, m+n)
	itemG := make([][]int, n)
	itemDeg := make([]int, n)
	for cur, items := range beforeItems {
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID { // 不同组项目，确定组间依赖关系
				groupG[preGroupID] = append(groupG[preGroupID], curGroupID)
				groupDeg[curGroupID]++
			} else { // 同组项目，确定组内依赖关系
				itemG[pre] = append(itemG[pre], cur)
				itemDeg[cur]++
			}
		}
	}

	// 组间拓扑序
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrders := topSort(groupG, groupDeg, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	// 按照组间的拓扑序，依次求得各个组的组内拓扑序，构成答案
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topSort(itemG, itemDeg, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return
}
