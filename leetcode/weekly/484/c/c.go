package main

// https://space.bilibili.com/206214
func countPairs1(words []string) (ans int64) {
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

func countPairs(words []string) int64 {
	cntWords := map[string]int{}
	for _, s := range words {
		cntWords[s]++
	}

	ans := 0
	cnt := map[string]int{}
	for s, c := range cntWords {
		t := []byte(s)
		base := t[0]
		for i := range t {
			t[i] = (t[i] - base + 26) % 26 // 保证结果在 [0, 25] 中
		}
		s = string(t)
		ans += cnt[s]*c + c*(c-1)/2 // c 个 s 中选 2 个有 C(c, 2) 种方案
		cnt[s] += c
	}
	return int64(ans)
}
