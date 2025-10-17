package main

// github.com/EndlessCheng/codeforces-go
func minCost(colors string, neededTime []int) (ans int) {
	maxT := 0
	for i, t := range neededTime {
		ans += t
		maxT = max(maxT, t)
		if i == len(colors)-1 || colors[i] != colors[i+1] {
			// 遍历到了连续同色段的末尾
			ans -= maxT // 不移除耗时最大的气球
			maxT = 0    // 准备计算下一段的最大耗时
		}
	}
	return
}
