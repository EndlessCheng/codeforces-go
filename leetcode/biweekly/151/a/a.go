package main

// https://space.bilibili.com/206214
func transformArray2(nums []int) []int {
	cnt1 := 0
	for _, x := range nums {
		cnt1 += x % 2
	}
	n := len(nums)
	cnt0 := n - cnt1
	clear(nums[:cnt0])
	for i := cnt0; i < n; i++ {
		nums[i] = 1
	}
	return nums
}

func transformArray(nums []int) []int {
	cnt := [2]int{}
	for _, x := range nums {
		cnt[x%2]++
	}
	clear(nums[:cnt[0]])
	for i := cnt[0]; i < len(nums); i++ {
		nums[i] = 1
	}
	return nums
}
