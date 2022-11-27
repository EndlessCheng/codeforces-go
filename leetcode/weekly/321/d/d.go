package main

// https://space.bilibili.com/206214
func countSubarrays(nums []int, k int) int {
	pos := 0
	for nums[pos] != k {
		pos++
	}

	cnt, c := map[int]int{0: 1}, 0 // i=pos 的时候 c 是 0，直接记到 cnt 中
	for _, x := range nums[pos+1:] {
		if x > k {
			c++
		} else {
			c--
		}
		cnt[c]++
	}

	ans := cnt[0] + cnt[1] // 子数组长为 1 和 2 的情况
	for i, c := pos-1, 0; i >= 0; i-- {
		if nums[i] < k {
			c++
		} else {
			c--
		}
		ans += cnt[c] + cnt[c+1]
	}
	return ans
}
