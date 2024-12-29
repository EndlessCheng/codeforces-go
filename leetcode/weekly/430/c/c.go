package main

// https://space.bilibili.com/206214
func numberOfSubsequences(nums []int) (ans int64) {
	n := len(nums)
	cnt := map[float32]int{}
	// 枚举 b 和 c
	for i := 4; i < n-2; i++ {
		// 增量式更新，本轮循环只需枚举 b=nums[i-2] 这一个数
		// 至于更前面的 b，已经在前面的循环中添加到 cnt 中了，不能重复添加
		b := float32(nums[i-2])
		// 枚举 a
		for _, a := range nums[:i-3] {
			cnt[float32(a)/b]++
		}

		c := float32(nums[i])
		// 枚举 d
		for _, d := range nums[i+2:] {
			ans += int64(cnt[float32(d)/c])
		}
	}
	return
}
