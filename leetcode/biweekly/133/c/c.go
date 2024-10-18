package main

// https://space.bilibili.com/206214
func minOperations(nums []int) int {
	ans := nums[0] ^ 1
	for i := 1; i < len(nums); i++ {
		ans += nums[i-1] ^ nums[i]
	}
	return ans
}

func minOperations2(nums []int) (ans int) {
	for _, x := range nums {
		if x == ans%2 {
			ans++
		}
	}
	return
}
