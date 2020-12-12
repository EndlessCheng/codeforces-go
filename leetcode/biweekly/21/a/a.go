package main

// github.com/EndlessCheng/codeforces-go
func sortString(s string) string {
	c := [26]int{}
	for _, b := range s {
		c[b-'a']++
	}
	ans := make([]byte, 0, len(s))
	for len(ans) < len(s) {
		for i := 0; i < 26; i++ {
			if c[i] > 0 {
				ans = append(ans, byte('a'+i))
				c[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if c[i] > 0 {
				ans = append(ans, byte('a'+i))
				c[i]--
			}
		}
	}
	return string(ans)
}
