package main

// https://space.bilibili.com/206214
func mirrorFrequency(s string) (ans int) {
	cnt := ['z' + 1]int{}
	for _, ch := range s {
		cnt[ch]++
	}

	for i := range 13 {
		ans += abs(cnt['a'+i] - cnt['z'-i])
	}
	for i := range 5 {
		ans += abs(cnt['0'+i] - cnt['9'-i])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
