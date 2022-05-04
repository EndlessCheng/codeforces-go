package main

// github.com/EndlessCheng/codeforces-go
func removeDigit(number string, digit byte) (ans string) {
	for i, ch := range number {
		if byte(ch) == digit {
			s := number[:i] + number[i+1:]
			if s > ans {
				ans = s
			}
		}
	}
	return
}
