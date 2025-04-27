package main

// https://space.bilibili.com/206214
func findCommonResponse(responses [][]string) (ans string) {
	maxCnt := 0
	cnt := map[string]int{}
	vis := map[string]struct{}{}
	for _, resp := range responses {
		clear(vis)
		for _, s := range resp {
			if _, ok := vis[s]; ok {
				continue
			}
			vis[s] = struct{}{}
			cnt[s]++
			c := cnt[s]
			if c > maxCnt || c == maxCnt && s < ans {
				maxCnt = c
				ans = s
			}
		}
	}
	return
}
