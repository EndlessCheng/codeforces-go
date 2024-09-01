package main

// https://space.bilibili.com/206214
func stringHash(s string, k int) string {
	n := len(s)
	ans := make([]byte, n/k)
	for i := 0; i < n; i += k {
		sum := 0
		for _, b := range s[i : i+k] {
			sum += int(b - 'a')
		}
		ans[i/k] = 'a' + byte(sum%26)
	}
	return string(ans)
}
