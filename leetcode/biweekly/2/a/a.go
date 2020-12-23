package main

// github.com/EndlessCheng/codeforces-go
func sumOfDigits(a []int) (ans int) {
	mi := int(1e9)
	for _, v := range a {
		if v < mi {
			mi = v
		}
	}
	for x := mi; x > 0; x /= 10 {
		ans ^= x % 10
	}
	return ans&1 ^ 1
}
