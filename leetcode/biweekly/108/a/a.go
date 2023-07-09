package main

// https://space.bilibili.com/206214
func alternatingSubarray(nums []int) int {
	ans := -1
	for i, n := 0, len(nums); i < n-1; {
		if nums[i+1]-nums[i] != 1 {
			i++
			continue
		}
		st := i
		for i++; i < n && nums[i] == nums[st]+(i-st)%2; i++ {
		}
		ans = max(ans, i-st)
		i--
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
