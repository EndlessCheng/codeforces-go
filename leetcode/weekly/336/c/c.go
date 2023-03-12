package main

// https://space.bilibili.com/206214
func beautifulSubarrays(nums []int) (ans int64) {
	s := 0
	cnt := map[int]int{s: 1}
	for _, x := range nums {
		s ^= x
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}
