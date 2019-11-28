package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(words)
	wScores := make([]int, n)
	//cnts := make([][26]int, n)
	for i, w := range words {
		for _, c := range w {
			wScores[i] += score[c-'a']
			//cnts[i][c-'a']++
		}
	}
	cnt := make([]int, 26)
	for _, l := range letters {
		cnt[l-'a']++
	}
	ans := 0
outer:
	for i := 0; i < (1 << uint(n)); i++ {
		cnt2 := make([]int, 26)
		sum := 0
		copy(cnt2, cnt)
		for j := 0; j < n; j++ {
			if i>>uint(j)&1 == 1 {
				for _, c := range words[j] {
					cnt2[c-'a']--
					if cnt2[c-'a'] < 0 {
						continue outer
					}
				}
				sum += wScores[j]
			}
		}
		ans = max(ans, sum)
	}
	return ans
}
