package main

// https://space.bilibili.com/206214/dynamic
func distinctNames(ideas []string) (ans int64) {
	size := [26]int{} // 集合大小
	intersection := [26][26]int{} // 交集大小
	groups := map[string]int{} // 后缀 -> 首字母
	for _, s := range ideas {
		b := s[0] - 'a'
		size[b]++ // 增加集合大小
		suffix := s[1:]
		mask := groups[suffix]
		groups[suffix] = mask | 1<<b // 把 b 加到 mask 中
		for a := 0; a < 26; a++ { // a 是和 s 有着相同后缀的首字母
			if mask>>a&1 > 0 { // a 在 mask 中
				intersection[b][a]++ // 增加交集大小
				intersection[a][b]++
			}
		}
	}

	for a := 1; a < 26; a++ { // 枚举所有组对
		for b := 0; b < a; b++ {
			m := intersection[a][b]
			ans += int64(size[a]-m) * int64(size[b]-m)
		}
	}
	return ans * 2
}

func distinctNames1(ideas []string) (ans int64) {
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
