package main

func numberOfSubarrays(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	cnt := make([]int, len(nums)+5)
	cnt[0]++
	for i, v := range nums {
		if v&1 == 1 {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i]
		}
		cnt[sum[i+1]]++
	}
	ans := 0
	for _, s := range sum {
		if s >= k {
			ans += cnt[s-k]
		}
	}
	return ans
}
