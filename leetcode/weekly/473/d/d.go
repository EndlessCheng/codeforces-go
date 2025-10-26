package main

// https://space.bilibili.com/206214
func numGoodSubarrays(nums []int, k int) (ans int64) {
	cnt := map[int]int{0: 1}
	sum := 0       // 前缀和
	lastStart := 0 // 上一个连续相同段的起始下标
	for i, x := range nums {
		if i > 0 && x != nums[i-1] {
			// 上一个连续相同段结束，可以把上一段对应的前缀和添加到 cnt
			s := sum
			for range i - lastStart {
				cnt[s%k]++
				s -= nums[i-1]
			}
			lastStart = i
		}
		sum += x
		ans += int64(cnt[sum%k])
	}
	return
}

func numGoodSubarrays1(nums []int, k int) (ans int64) {
	n := len(nums)
	cnt := map[int]int{0: 1}
	s := 0
	for i := 0; i < n; {
		start := i
		x := nums[start]
		s0 := s
		for ; i < n && nums[i] == x; i++ {
			s += x
			ans += int64(cnt[s%k])
		}
		for range i - start {
			s0 += x
			cnt[s0%k]++
		}
	}
	return
}
