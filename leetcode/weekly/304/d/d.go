package main

// https://space.bilibili.com/206214
func longestCycle(edges []int) int {
	time := make([]int, len(edges))
	clock, ans := 1, -1
	for x, t := range time {
		if t > 0 {
			continue
		}
		for startTime := clock; x >= 0; x = edges[x] {
			if time[x] > 0 { // 重复访问
				if time[x] >= startTime { // 找到了一个新的环
					ans = max(ans, clock-time[x])
				}
				break
			}
			time[x] = clock
			clock++
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
