package main

// https://space.bilibili.com/206214
func maximumRobots(chargeTimes, runningCosts []int, budget int64) (ans int) {
	sum, left, q := int64(0), 0, []int{}
	// 枚举区间右端点 right，计算区间左端点 left 的最小值
	for right, t := range chargeTimes {
		// 及时清除队列中的无用数据，保证队列的单调性
		for len(q) > 0 && t >= chargeTimes[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, right)
		sum += int64(runningCosts[right])
		// 如果左端点 left 不满足要求，就不断右移 left
		for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*sum > budget {
			// 及时清除队列中的无用数据，保证队列的单调性
			if q[0] == left {
				q = q[1:]
			}
			sum -= int64(runningCosts[left])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
