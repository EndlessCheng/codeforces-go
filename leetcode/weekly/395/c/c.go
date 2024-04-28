package main

// https://space.bilibili.com/206214
func minEnd(n, x int) int64 {
	n-- // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
	j := 0
	for t, lb := ^x, 0; n>>j > 0; t ^= lb {
		lb = t & -t
		x |= n >> j & 1 * lb
		j++
	}
	return int64(x)
}

func minEnd2(n, x int) int64 {
	n-- // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
	i, j := 0, 0
	for n>>j > 0 {
		// x 的第 i 个比特值是 0，即「空位」
		if x>>i&1 == 0 {
			// 空位填入 n 的第 j 个比特值
			x |= n >> j & 1 << i
			j++
		}
		i++
	}
	return int64(x)
}
