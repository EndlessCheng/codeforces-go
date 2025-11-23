package main

// https://space.bilibili.com/206214
func maxBalancedSubarray(nums []int) (ans int) {
	type pair struct{ xor, diff int }
	pos := map[pair]int{{}: -1}
	xor, diff := 0, 0
	for i, x := range nums {
		xor ^= x
		diff += x%2*2 - 1
		p := pair{xor, diff}
		if j, ok := pos[p]; ok {
			ans = max(ans, i-j)
		} else {
			pos[p] = i
		}
	}
	return
}
