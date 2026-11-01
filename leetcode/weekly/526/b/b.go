package main

// https://space.bilibili.com/206214
// 返回 value 的最小质因子
func minPrimeDivisor(value int) int {
	for p := 2; p*p <= value; p++ {
		if value%p == 0 {
			return p
		}
	}
	return value // value 是质数或者 1
}

func minimumMaxValue(n, index, value int) int64 {
	minP := minPrimeDivisor(value)
	if value == 1 || value-minP-(index-1)*2 <= 0 { // nums[0] <= 0
		return -1
	}
	if index == n-1 {
		return int64(value)
	}
	return int64(value + minP + (n-2-index)*2)
}
