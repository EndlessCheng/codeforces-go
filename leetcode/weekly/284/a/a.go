package main

// github.com/EndlessCheng/codeforces-go
func findKDistantIndices1(nums []int, key, k int) (ans []int) {
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

func findKDistantIndices(nums []int, key, k int) (ans []int) {
	j := 0
	for i, x := range nums {
		if x != key {
			continue
		}
		j = max(j, i-k) // j 至少是 i-k
		for j <= min(i+k, len(nums)-1) { // j 至多是 i+k，但不能越界
			ans = append(ans, j)
			j++
		}
	}
	return
}
