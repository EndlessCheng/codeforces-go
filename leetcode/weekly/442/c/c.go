package main

// https://space.bilibili.com/206214
func minTime(skill, mana []int) int64 {
	n := len(skill)
	lastFinish := make([]int, n) // 第 i 个巫师完成上一瓶药水的时间
	for _, m := range mana {
		// 按题意模拟
		sumT := 0
		for i, x := range skill {
			sumT = max(sumT, lastFinish[i]) + x*m
		}
		// 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
		lastFinish[n-1] = sumT
		for i := n - 2; i >= 0; i-- {
			lastFinish[i] = lastFinish[i+1] - skill[i+1]*m
		}
	}
	return int64(lastFinish[n-1])
}
