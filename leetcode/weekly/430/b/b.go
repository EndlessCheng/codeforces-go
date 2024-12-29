package main

// https://space.bilibili.com/206214
func answerString(s string, k int) (ans string) {
	if k == 1 {
		return s
	}
	n := len(s)
	for i := range n {
		ans = max(ans, s[i:min(i+n-k+1, n)])
	}
	return
}
