package main

// https://space.bilibili.com/206214/dynamic
func rearrangeCharacters(s, target string) int {
	var cntS, cntT [26]int
	for _, ch := range s { cntS[ch-'a']++ }
	for _, ch := range target { cntT[ch-'a']++ }
	ans := len(s)
	for i, c := range cntS {
		if cntT[i] > 0 {
			ans = min(ans, c/cntT[i])
		}
	}
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
