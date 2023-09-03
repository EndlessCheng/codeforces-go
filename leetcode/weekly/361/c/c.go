package main

// https://space.bilibili.com/206214
func countInterestingSubarrays(nums []int, mod, k int) (ans int64) {
	cnt := map[int]int{0: 1} // 把 s[0]=0 算进去
	s := 0
	for _, x := range nums {
		if x%mod == k {
			s = (s + 1) % mod
		}
		ans += int64(cnt[(s-k+mod)%mod])
		cnt[s]++
	}
	return
}
