package main

// https://space.bilibili.com/206214
func taskSchedulerII(tasks []int, space int) int64 {
	ans := 0
	last := map[int]int{}
	for _, t := range tasks {
		ans++ // 完成该任务，天数+1
		if last[t] > 0 {
			ans = max(ans, last[t]+space+1) // 看看是否要间隔 space 天
		}
		last[t] = ans // 记录上一次完成时间
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
