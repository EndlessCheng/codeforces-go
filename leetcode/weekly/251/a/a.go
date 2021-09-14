package main

// Go 模拟

// github.com/EndlessCheng/codeforces-go
func getLucky(s string, k int) (sum int) {
	for _, b := range s {
		b -= 'a' - 1
		sum += int(b/10 + b%10)
	}
	for k--; k > 0 && sum > 9; k-- {
		s := sum
		for sum = 0; s > 0; s /= 10 {
			sum += s % 10
		}
	}
	return
}
