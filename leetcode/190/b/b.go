package main

func maxVowels(s string, k int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)

	c := ['z' + 1]int{}
	for _, b := range s[:k] {
		c[b]++
	}
	for _, b := range "aeiou" {
		ans += c[b]
	}
	i, j := 0, k
	for ; j < n; j++ {
		c[s[i]]--
		c[s[j]]++
		s := 0
		for _, b := range "aeiou" {
			s += c[b]
		}
		ans = max(ans, s)
		i++
	}
	return
}
