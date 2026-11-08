package main

// https://space.bilibili.com/206214
func repeatedStringMatch(a, b string, x, y int) bool {
	n, m := len(a), len(b)
	lenA, lenB := n*x, m*y
	if lenB > lenA {
		return false
	}

	firsts := min(lcm(n, m), lenB)
next:
	for start := range min(n, lenA-lenB+1) {
		// 判断 B 是否为 A[start:] 的前缀
		for i := range firsts { // 只需比较前 firsts 个字母
			if a[(start+i)%n] != b[i%m] {
				continue next
			}
		}
		return true
	}
	return false
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
