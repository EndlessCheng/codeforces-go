package main

// https://space.bilibili.com/206214
func maximumSubarraySum(nums []int, k int) (ans int64) {
	s := int64(0)
	cnt := map[int]int{}
	for i, x := range nums {
		// 1. 进入窗口
		s += int64(x)
		cnt[x]++

		left := i - k + 1
		if left < 0 { // 窗口大小不足 k
			continue
		}

		// 2. 更新答案
		if len(cnt) == k {
			ans = max(ans, s)
		}

		// 3. 离开窗口
		out := nums[left]
		s -= int64(out)
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return
}
