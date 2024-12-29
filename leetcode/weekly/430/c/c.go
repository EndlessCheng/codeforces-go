package main

// https://space.bilibili.com/206214
func numberOfSubsequences(nums []int) (ans int64) {
	n := len(nums)
	cnt := map[float64]int{}
	// 枚举 b 和 c
	for i := 4; i < n-2; i++ {
		b := float64(nums[i-2])
		// 枚举 a
		for _, a := range nums[:i-3] {
			cnt[float64(a)/b]++
		}

		c := float64(nums[i])
		// 枚举 d
		for _, d := range nums[i+2:] {
			ans += int64(cnt[float64(d)/c])
		}
	}
	return
}
