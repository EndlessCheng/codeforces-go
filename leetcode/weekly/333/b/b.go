package main

// https://space.bilibili.com/206214
func minOperations(n int) int {
	ans := 1
	for n&(n-1) > 0 { // n 不是 2 的幂次
		lb := n & -n
		if n&(lb<<1) > 0 { // 多个连续 1
			n += lb
		} else {
			n -= lb // 单个 1
		}
		ans++
	}
	return ans
}
