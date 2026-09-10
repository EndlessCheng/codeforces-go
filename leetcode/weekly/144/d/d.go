package main

// https://space.bilibili.com/206214
func maxDepthAfterSplit(seq string) []int {
	ans := make([]int, len(seq))
	for i, ch := range seq {
		ans[i] = (i + int(ch)) % 2
	}
	return ans
}
