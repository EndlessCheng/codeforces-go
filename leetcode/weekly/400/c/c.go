package main

// https://space.bilibili.com/206214
func clearStars(S string) string {
	s := []byte(S)
	st := make([][]int, 26)
	for i, c := range s {
		if c != '*' {
			st[c-'a'] = append(st[c-'a'], i)
			continue
		}
		for j, ps := range st {
			if m := len(ps); m > 0 {
				s[ps[m-1]] = '*'
				st[j] = ps[:m-1]
				break
			}
		}
	}

	t := s[:0]
	for _, c := range s {
		if c != '*' {
			t = append(t, c)
		}
	}
	return string(t)
}
