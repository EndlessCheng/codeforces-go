package c2

const mx = 5

var c [mx][mx]int

func init() {
	for i := range mx {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

// 计算 C(n, k) % p，要求 p 是质数
func lucas(n, k, p int) int {
	if k == 0 {
		return 1
	}
	return c[n%p][k%p] * lucas(n/p, k/p, p) % p
}

func comb(n, k int) int {
	// 结果至多为 5 + 4 * 6 = 29，无需中途取模
	return lucas(n, k, 2)*5 + lucas(n, k, 5)*6
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%10 == 0
}
