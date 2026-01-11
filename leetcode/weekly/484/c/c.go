package main

// https://space.bilibili.com/206214
func countPairs(words []string) (ans int64) {
	cnt := map[string]int{}
	for _, s := range words {
		t := []byte(s)
		base := t[0]
		for i := range t {
			t[i] = (t[i] - base + 26) % 26 // 保证结果在 [0, 25] 中
		}
		s = string(t)
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}
