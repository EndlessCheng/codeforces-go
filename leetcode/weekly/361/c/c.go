package main

// https://space.bilibili.com/206214
func countInterestingSubarrays(nums []int, modulo, k int) (ans int64) {
	cnt := make([]int, min(len(nums)+1, modulo))
	cnt[0] = 1 // 单独统计 s[0]=0
	s := 0
	for _, x := range nums {
		if x%modulo == k {
			s++
		}
		if s >= k {
			ans += int64(cnt[(s-k)%modulo])
		}
		cnt[s%modulo]++
	}
	return
}

func countInterestingSubarrays1(nums []int, modulo, k int) (ans int64) {
	sum := make([]int, len(nums)+1)
	for i, x := range nums {
		sum[i+1] = sum[i]
		if x%modulo == k {
			sum[i+1]++
		}
	}

	cnt := make([]int, min(len(nums)+1, modulo))
	for _, s := range sum {
		if s >= k {
			ans += int64(cnt[(s-k)%modulo])
		}
		cnt[s%modulo]++
	}
	return
}
