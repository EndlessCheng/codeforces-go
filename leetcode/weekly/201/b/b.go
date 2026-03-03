package main

// github.com/EndlessCheng/codeforces-go
func findKthBit1(n, k int) byte {
	if n == 1 {
		return '0'
	}
	if k == 1<<(n-1) {
		return '1'
	}
	if k < 1<<(n-1) {
		return findKthBit(n-1, k)
	}
	return findKthBit(n-1, 1<<n-k) ^ 1
}

func findKthBit2(n, k int) byte {
	rev := byte(0) // 翻转次数的奇偶性
	for {
		if n == 1 {
			return '0' ^ rev
		}
		if k == 1<<(n-1) {
			return '1' ^ rev
		}
		if k > 1<<(n-1) {
			k = 1<<n - k
			rev ^= 1
		}
		n--
	}
}

func findKthBit(_, k int) byte {
	if k%2 > 0 {
		return '0' + byte(k/2%2)
	}
	k /= k & -k // 去掉 k 的尾零
	return '1' - byte(k/2%2)
}
