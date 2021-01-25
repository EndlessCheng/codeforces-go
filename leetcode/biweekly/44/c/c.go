package main

// github.com/EndlessCheng/codeforces-go
func decode(encoded []int) (ans []int) {
	n := len(encoded)
	a0 := 0
	for i := 1; i <= n+1; i++ {
		a0 ^= i
	}
	for i := 1; i < n; i += 2 {
		a0 ^= encoded[i]
	}
	ans = append(ans, a0)
	for _, v := range encoded {
		ans = append(ans, ans[len(ans)-1]^v)
	}
	return
}
