package main

// https://space.bilibili.com/206214
func splitArray(nums []int) int64 {
	n := len(nums)
	// 最长递增前缀
	pre := nums[0]
	i := 1
	for i < n && nums[i] > nums[i-1] {
		pre += nums[i]
		i++
	}

	// 最长递增后缀
	suf := nums[n-1]
	j := n - 2
	for j >= 0 && nums[j] > nums[j+1] {
		suf += nums[j]
		j--
	}

	// 情况一
	if i-1 < j {
		return -1
	}

	d := pre - suf
	// 情况二
	if i-1 == j {
		return int64(abs(d))
	}

	// 情况三
	return int64(min(abs(d+nums[i-1]), abs(d-nums[i-1])))
}

func abs(x int) int { if x < 0 { return -x }; return x }
