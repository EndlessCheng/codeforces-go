package main

// https://space.bilibili.com/206214
func smallestAbsent(nums []int) int {
	has := map[int]bool{}
	sum := 0
	for _, x := range nums {
		has[x] = true
		sum += x
	}

	ans := max(sum/len(nums)+1, 1) // 答案必须是正整数
	for has[ans] {
		ans++
	}
	return ans
}
