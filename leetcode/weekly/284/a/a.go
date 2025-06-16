package main

// github.com/EndlessCheng/codeforces-go
func findKDistantIndices(nums []int, key, k int) (ans []int) {
	last := -k - 1
	for i := k - 1; i >= 0; i-- {
		if nums[i] == key {
			last = i
			break
		}
	}

	for i := range nums {
		if i+k < len(nums) && nums[i+k] == key {
			last = i + k
		}
		if last >= i-k {
			ans = append(ans, i)
		}
	}
	return
}
