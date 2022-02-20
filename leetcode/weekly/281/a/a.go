package main

// github.com/EndlessCheng/codeforces-go
func countEven(n int) (ans int) {
	for i := 1; i <= n; i++ {
		s := 0
		for x := i; x > 0; x /= 10 {
			s ^= x % 10 & 1
		}
		ans += s ^ 1 // 偶数就 +1
	}
	return
}
