package main

// https://space.bilibili.com/206214
func countValidPrefixes(s string) (ans int) {
	cnt := [2]int{}
	for _, ch := range s {
		cnt[ch-'0']++
		if abs(cnt[0]-cnt[1]) <= 1 {
			ans++
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
