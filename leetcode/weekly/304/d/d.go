package main

// https://space.bilibili.com/206214
func longestCycle(edges []int) int {
	ans := -1
	visTime := make([]int, len(edges)) // 首次访问 x 的时间
	curTime := 1 // 当前时间
	for x := range edges {
		startTime := curTime // 本轮循环的开始时间
		for x != -1 && visTime[x] == 0 { // 没有访问过 x
			visTime[x] = curTime // 记录访问 x 的时间
			curTime++
			x = edges[x] // 访问下一个节点
		}
		// 如果 0 < visTime[x] < startTime，说明 x 在本轮循环开始之前就访问过，即使 x 在环上或者可以到达环，我们也已经更新了 ans
		// 如果 visTime[x] >= startTime，说明 x 是在本轮循环中访问的点，我们找到了一个环
		if x != -1 && visTime[x] >= startTime {
			ans = max(ans, curTime-visTime[x]) // 当前时间减去上次访问 x 的时间，即为环长
		}
	}
	return ans // 注意，如果没有找到环，返回的是 ans 的初始值 -1
}
