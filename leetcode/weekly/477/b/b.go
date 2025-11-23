package main

// https://space.bilibili.com/206214
func maxBalancedSubarray(nums []int) (ans int) {
	type pair struct{ xor, diff int }
	pos := map[pair]int{{}: -1} // 空前缀的位置视作 -1
	p := pair{}
	for i, x := range nums {
		p.xor ^= x
		p.diff += x%2*2 - 1
		if j, ok := pos[p]; ok {
			ans = max(ans, i-j)
		} else {
			pos[p] = i
		}
	}
	return
}
