package main

// https://space.bilibili.com/206214
func maxDivScore(nums, divisors []int) (ans int) {
	maxCnt := -1
	for _, d := range divisors {
		cnt := 0
		for _, x := range nums {
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt || cnt == maxCnt && d < ans {
			maxCnt, ans = cnt, d
		}
	}
	return
}
