package main

// https://space.bilibili.com/206214
func countCompleteSubarrays(nums []int) (ans int) {
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}
	k := len(set)

	cnt := make(map[int]int, k) // 预分配空间
	left := 0
	for _, x := range nums {
		cnt[x]++
		for len(cnt) == k {
			out := nums[left]
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
			left++
		}
		ans += left
	}
	return
}
