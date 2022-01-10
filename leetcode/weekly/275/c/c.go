package main

// github.com/EndlessCheng/codeforces-go
func wordCount(startWords, targetWords []string) (ans int) {
	has := map[int]bool{}
	for _, word := range startWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		has[mask] = true
	}
	for _, word := range targetWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		for i := 0; i < 26; i++ {
			if mask&(1<<i) > 0 && has[mask^(1<<i)] { // 去掉这个字符
				ans++
				break
			}
		}
	}
	return
}
