package main

// https://space.bilibili.com/206214/dynamic
func distinctNames(ideas []string) (ans int64) {
	group := [26]map[string]bool{}
	for i := range group {
		group[i] = map[string]bool{}
	}
	for _, s := range ideas {
		group[s[0]-'a'][s[1:]] = true
	}
	for i, a := range group {
		for _, b := range group[:i] {
			m := 0
			for s := range a {
				if b[s] {
					m++
				}
			}
			ans += int64(len(a)-m) * int64(len(b)-m)
		}
	}
	return ans * 2
}

func distinctNames2(ideas []string) (ans int64) {
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
