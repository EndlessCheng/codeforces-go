package main

// https://space.bilibili.com/206214
func numberOfPairs(nums1, nums2 []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums1 {
		if x%k > 0 {
			continue
		}
		x /= k
		for d := 1; d*d <= x; d++ {
			if x%d == 0 {
				cnt[d]++
				if d*d < x {
					cnt[x/d]++
				}
			}
		}
	}
	for _, x := range nums2 {
		ans += cnt[x]
	}
	return
}
