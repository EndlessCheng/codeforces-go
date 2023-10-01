package main

// https://space.bilibili.com/206214
func maxSubarrays(nums []int) (ans int) {
	and := -1 // -1 就是 111...1，和任何数 AND 都等于那个数
	for _, x := range nums {
		and &= x
		if and == 0 {
			ans++
			and = -1
		}
	}
	return max(ans, 1) // 如果 ans=0 说明所有数的 and>0，答案为 1
}

func max(a, b int) int { if b > a { return b }; return a }
