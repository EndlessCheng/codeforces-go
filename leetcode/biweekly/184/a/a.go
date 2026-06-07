package main

// https://space.bilibili.com/206214
func consecutiveSetBits1(n int) bool {
	cnt := 0
	for ; n > 1; n >>= 1 {
		if n&3 == 3 { // 最低两位是 11
			cnt++
		}
	}
	return cnt == 1
}

func consecutiveSetBits(n int) bool {
	m := n & (n >> 1)
	return m > 0 && m&(m-1) == 0
}
