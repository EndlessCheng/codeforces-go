package main

// https://space.bilibili.com/206214
func clearStars(s string) string {
	st := [26][]int{}
	del := make([]bool, len(s))
	for i, c := range s {
		if c != '*' {
			st[c-'a'] = append(st[c-'a'], i)
			continue
		}
		for j, ps := range st {
			if m := len(ps); m > 0 {
				del[ps[m-1]] = true
				st[j] = ps[:m-1]
				break
			}
		}
	}

	t := []byte{}
	for i, d := range del {
		if !d && s[i] != '*' {
			t = append(t, s[i])
		}
	}
	return string(t)
}
