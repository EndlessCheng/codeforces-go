package main

// github.com/EndlessCheng/codeforces-go
func getDescentPeriods(prices []int) (ans int64) {
	dec := 0
	for i, p := range prices {
		if i > 0 && p == prices[i-1]-1 {
			dec++ // 连续递减
		} else {
			dec = 1 // 连续递减中断，重新统计
		}
		ans += int64(dec) // dec 是右端点为 i 的连续递减子数组个数
	}
	return
}
