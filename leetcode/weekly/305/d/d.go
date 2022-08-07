package main

// https://space.bilibili.com/206214
func longestIdealString(s string, k int) (ans int) {
	f := [26]int{}
	for _, c := range s {
		c := int(c - 'a')
		for j := max(c-k, 0); j <= min(c+k, 25); j++ {
			f[c] = max(f[c], f[j])
		}
		f[c]++
	}
	for _, v := range f {
		ans = max(ans, v)
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
