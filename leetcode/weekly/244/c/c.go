package main

// github.com/EndlessCheng/codeforces-go
func minFlips(s string) int {
	n := len(s)
	ans := n
	cnt := 0
	for i := range n*2 - 1 {
		cnt += (int(s[i%n]) ^ i) & 1
		left := i - n + 1
		if left < 0 {
			continue
		}
		ans = min(ans, cnt, n-cnt)
		cnt -= (int(s[left]) ^ left) & 1
	}
	return ans
}
