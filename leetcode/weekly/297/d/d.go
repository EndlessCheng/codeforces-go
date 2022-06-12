package main

// https://space.bilibili.com/206214/dynamic
func distinctNames(ideas []string) (ans int64) {
	group := map[string]int{}
	for _, s := range ideas {
		group[s[1:]] |= 1 << (s[0] - 'a')
	}
	cnt := [26][26]int{}
	for _, mask := range group {
		for i := 0; i < 26; i++ {
			if mask>>i&1 == 0 {
				for j := 0; j < 26; j++ {
					if mask>>j&1 > 0 {
						cnt[i][j]++
					}
				}
			} else {
				for j := 0; j < 26; j++ {
					if mask>>j&1 == 0 {
						ans += int64(cnt[i][j])
					}
				}
			}
		}
	}
	return ans * 2
}
