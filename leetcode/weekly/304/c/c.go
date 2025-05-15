package main

// https://space.bilibili.com/206214
func closestMeetingNode(edges []int, node1, node2 int) int {
	n := len(edges)
	calcDis := func(x int) []int {
		dis := make([]int, n)
		for i := range dis {
			dis[i] = n
		}
		// 从 x 出发，直到无路可走（x=-1）或者重复访问节点（dis[x]<n）
		for d := 0; x >= 0 && dis[x] == n; x = edges[x] {
			dis[x] = d
			d++
		}
		return dis
	}

	dis1 := calcDis(node1)
	dis2 := calcDis(node2)

	minDis, ans := n, -1
	for i, d1 := range dis1 {
		d := max(d1, dis2[i])
		if d < minDis {
			minDis, ans = d, i
		}
	}
	return ans
}
