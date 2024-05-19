package main

import "strconv"

// https://space.bilibili.com/206214
func sumDigitDifferences(nums []int) int64 {
	n, m := len(nums), len(strconv.Itoa(nums[0]))
	ans := m * n * (n - 1) / 2
	cnt := make([][10]int, m)
	for _, x := range nums {
		for i := 0; x > 0; x /= 10 {
			d := x % 10
			ans -= cnt[i][d]
			cnt[i][d]++
			i++
		}
	}
	return int64(ans)
}
