package main

// https://space.bilibili.com/206214
func countSubarrays(nums []int, k int) (ans int) {
	cnt := map[int]int{0: 1} // k 左边的前缀和的出现次数
	sum := 0
	foundK := false
	for _, x := range nums {
		if x == k {
			foundK = true
		} else if x < k {
			sum--
		} else {
			sum++
		}
		if !foundK {
			cnt[sum]++ // 统计 k 左边的前缀和的出现次数
		} else {
			ans += cnt[sum] + cnt[sum-1]
		}
	}
	return
}

func countSubarrays2(nums []int, k int) (ans int) {
	n := len(nums)
	cnt := make([]int, n*2)
	cnt[n] = 1
	sum := n
	foundK := false
	for _, x := range nums {
		if x == k {
			foundK = true
		} else if x < k {
			sum--
		} else {
			sum++
		}
		if !foundK {
			cnt[sum]++ // 统计 k 左边的前缀和的出现次数
		} else {
			ans += cnt[sum] + cnt[sum-1]
		}
	}
	return
}
