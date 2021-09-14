package main

// github.com/EndlessCheng/codeforces-go
func longestBeautifulSubstring(s string) (ans int) {
	const vowel = "aeiou"
	cur, sum := 0, 0
	for i, n := 0, len(s); i < n; {
		st := i
		v := s[st]
		for ; i < n && s[i] == v; i++ {
		}

		if v != vowel[cur] {
			cur, sum = 0, 0
			if v != vowel[0] {
				continue
			}
		}

		sum += i - st
		if cur++; cur == 5 {
			if sum > ans {
				ans = sum
			}
			cur, sum = 0, 0
		}
	}
	return
}
