package main

// github.com/EndlessCheng/codeforces-go
func minSteps(s, t string) (ans int) {
	cnt := [26]int{}
	for _, ch := range s { cnt[ch-'a']++ }
	for _, ch := range t { cnt[ch-'a']-- }
	for _, c := range cnt { ans += abs(c) }
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
