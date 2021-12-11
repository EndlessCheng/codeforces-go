package main

// 前后缀经典技巧

// github.com/EndlessCheng/codeforces-go
func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	if time*2 >= n { // i 最小为 time，而且需要满足 i+time<n，即 time*2<n，不满足则直接返回空数组
		return
	}

	suf := make([]int, n) // suf[i] 表示从 security[i+1] 向后的最长连续非递减长度
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			suf[i] = suf[i+1] + 1
		}
	}

	pre := 0 // pre 表示从 security[i-1] 向前的最长连续非递增长度
	for i, v := range security[:n-time] {
		if i > 0 && v <= security[i-1] {
			pre++
		} else {
			pre = 0
		}
		if i >= time && pre >= time && suf[i] >= time {
			ans = append(ans, i)
		}
	}
	return
}
