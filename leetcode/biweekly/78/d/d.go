package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func largestVariance2(s string) (ans int) {
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			if b == a {
				continue
			}
			diff, diffWithB := 0, -len(s)
			for _, ch := range s {
				if ch == a {
					diff++
					diffWithB++
				} else if ch == b {
					diff--
					diffWithB = diff // 记录包含 b 时的 diff
					diff = max(diff, 0)
				}
				ans = max(ans, diffWithB)
			}
		}
	}
	return
}

func largestVariance(s string) (ans int) {
	if strings.Count(s, s[:1]) == len(s) {
		return
	}
	var diff, diffWithB [26][26]int
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			diffWithB[i][j] = -len(s)
		}
	}
	for _, ch := range s {
		ch -= 'a'
		for i := rune(0); i < 26; i++ {
			if i == ch {
				continue
			}
			diff[ch][i]++ // a=ch, b=i
			diffWithB[ch][i]++
			diff[i][ch]-- // a=i, b=ch
			diffWithB[i][ch] = diff[i][ch]
			diff[i][ch] = max(diff[i][ch], 0)
			ans = max(ans, max(diffWithB[ch][i], diffWithB[i][ch]))
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
