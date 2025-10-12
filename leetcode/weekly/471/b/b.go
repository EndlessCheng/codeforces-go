package main

// https://space.bilibili.com/206214
func longestBalanced(s string) (ans int) {
	for i := range s {
		cnt := make([]int, 26)
	next:
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			base := 0
			for _, c := range cnt {
				if c == 0 {
					continue
				}
				if base == 0 {
					base = c
				} else if c != base {
					continue next
				}
			}
			ans = max(ans, j-i+1)
		}
	}
	return
}
