package main

// github.com/EndlessCheng/codeforces-go
func findNumOfValidWords(words []string, puzzles []string) (ans []int) {
	cnt := map[int]int{}
	for _, s := range words {
		v := 0
		for _, b := range s {
			v |= 1 << (b - 'a')
		}
		cnt[v]++
	}
	ans = make([]int, len(puzzles))
	for i, s := range puzzles {
		bit := 1 << (s[0] - 'a')
		v := 0
		for _, b := range s[1:] {
			v |= 1 << (b - 'a')
		}
		for sub, ok := v, true; ok; ok = sub != v {
			ans[i] += cnt[sub|bit]
			sub = (sub - 1) & v
		}
	}
	return
}
