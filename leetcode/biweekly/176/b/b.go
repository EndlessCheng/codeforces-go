package main

// https://space.bilibili.com/206214
func prefixConnected(words []string, k int) (ans int) {
	cnt := map[string]int{}
	for _, w := range words {
		if len(w) >= k {
			cnt[w[:k]]++
		}
	}

	for _, c := range cnt {
		if c > 1 {
			ans++
		}
	}
	return
}
