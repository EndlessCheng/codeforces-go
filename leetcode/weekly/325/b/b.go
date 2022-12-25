package main

// https://space.bilibili.com/206214
func takeCharacters(s string, k int) int {
	n := len(s)
	c, j := [3]int{}, n
	for c[0] < k || c[1] < k || c[2] < k {
		if j == 0 {
			return -1
		}
		j--
		c[s[j]-'a']++
	}
	ans := n - j
	for i := 0; i < n && j < n; i++ {
		c[s[i]-'a']++
		for j < n && c[s[j]-'a'] > k {
			c[s[j]-'a']--
			j++
		}
		ans = min(ans, i+1+n-j)
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
