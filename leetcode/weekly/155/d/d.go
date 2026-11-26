package main

// 模板来源：https://leetcode.cn/discuss/post/3581143/
// 返回有向图 g 的节点子集 subNodes 的拓扑序
// 如果图中有环，返回 nil
func topologicalSort(g [][]int, inDeg []int, subNodes []int) []int {
	n := len(subNodes)
	q := make([]int, 0, n)
	topoOrder := q
	for _, x := range subNodes {
		if inDeg[x] == 0 { // 节点 x 没有依赖，加入拓扑序
			q = append(q, x)
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		// 节点 y 依赖节点 x
		for _, y := range g[x] {
			inDeg[y]-- // y 的依赖数量减一
			if inDeg[y] == 0 { // y 的依赖都排在 y 的左边了，可以排 y
				q = append(q, y)
			}
		}
	}

	if cap(q) > 0 { // 有节点没入队，说明图中有环
		return nil
	}
	return topoOrder[:n]
}

func sortItems(n int, m int, group []int, beforeItems [][]int) (ans []int) {
	// groupItems[i] 保存组 i 的所有项目
	groupItems := make([][]int, m+n)
	for i, g := range group {
		if g < 0 { // 项目 i 不属于任何小组
			group[i] = m // 创建一个新的小组，方便拓扑排序
			g = m
			m++
		}
		groupItems[g] = append(groupItems[g], i)
	}

	// 建图，顺带计算入度
	groupG := make([][]int, m)
	groupInDeg := make([]int, m)
	itemG := make([][]int, n)
	itemInDeg := make([]int, n)
	for x, items := range beforeItems {
		gx := group[x]
		for _, y := range items {
			gy := group[y]
			if gx != gy {
				groupG[gy] = append(groupG[gy], gx) // 组与组的依赖关系
				groupInDeg[gx]++
			} else {
				itemG[y] = append(itemG[y], x) // 同组项目的依赖关系
				itemInDeg[x]++
			}
		}
	}

	// 计算组与组之间的顺序
	groupIDs := make([]int, m)
	for i := range groupIDs {
		groupIDs[i] = i
	}
	groupOrder := topologicalSort(groupG, groupInDeg, groupIDs)
	if groupOrder == nil { // 循环依赖
		return
	}

	// 对于每个组，计算组内的项目顺序
	for _, i := range groupOrder {
		itemOrder := topologicalSort(itemG, itemInDeg, groupItems[i])
		if itemOrder == nil { // 循环依赖
			return nil
		}
		ans = append(ans, itemOrder...)
	}
	return
}
