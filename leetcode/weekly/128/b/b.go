package main

// github.com/EndlessCheng/codeforces-go
func numPairsDivisibleBy60(time []int) (ans int) {
	cnt := [60]int{}
	for _, t := range time {
		// 先查询 cnt，再更新 cnt，因为题目要求 i<j
		// 如果先更新，再查询，就把 i=j 的情况也考虑进去了
		ans += cnt[(60-t%60)%60]
		cnt[t%60]++
	}
	return
}
