package main

// https://space.bilibili.com/206214
func similarPairs(words []string) (ans int) {
	cnt := map[int]int{}
	for _, s := range words {
		mask := 0
		for _, c := range s {
			mask |= 1 << (c - 'a')
		}
		ans += cnt[mask]
		cnt[mask]++
	}
	return
}
