package main

// github.com/EndlessCheng/codeforces-go
func kthDistinct(a []string, k int) (ans string) {
	cnt := map[string]int{}
	for _, s := range a {
		cnt[s]++
	}
	for _, s := range a {
		if cnt[s] == 1 {
			if k--; k == 0 {
				return s
			}
		}
	}
	return
}
