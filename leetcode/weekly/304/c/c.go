package main

// https://space.bilibili.com/206214
func closestMeetingNode(edges []int, x, y int) int {
	n := len(edges)
	ans := n
	visX := make([]bool, n)
	visY := make([]bool, n)

	for !visX[x] || !visY[y] { // x 或 y 没有访问过
		visX[x] = true // 标记访问过
		visY[y] = true

		if visY[x] { // 我吹过你吹过的晚风
			ans = x
		}
		if visX[y] {
			ans = min(ans, y) // 如果有多个答案，返回最小的节点编号
		}
		if ans < n {
			return ans
		}

		if edges[x] >= 0 {
			x = edges[x] // 继续走
		}
		if edges[y] >= 0 {
			y = edges[y] // 继续走
		}
	}

	return -1
}

func closestMeetingNode1(edges []int, node1, node2 int) int {
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
