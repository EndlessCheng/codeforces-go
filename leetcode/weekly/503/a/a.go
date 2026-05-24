package main

// https://space.bilibili.com/206214
func limitOccurrences(nums []int, k int) []int {
	ans := nums[:0]
	cnt := 0
	for i, x := range nums {
		if i == 0 || x == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= k {
			ans = append(ans, x)
		}
	}
	return ans
}

func limitOccurrences2(nums []int, k int) []int {
	cnt := 0
	j := 0
	for i, x := range nums {
		if i == 0 || x == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= k {
			nums[j] = x
			j++
		}
	}
	return nums[:j]
}
