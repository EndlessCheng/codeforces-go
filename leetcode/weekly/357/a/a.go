package main

// https://space.bilibili.com/206214
func finalString(s string) string {
	qs := [2][]rune{}
	dir := 1
	for _, c := range s {
		if c == 'i' {
			dir ^= 1
		} else {
			qs[dir] = append(qs[dir], c)
		}
	}
	q := qs[dir^1]
	for i, n := 0, len(q); i < n/2; i++ {
		q[i], q[n-1-i] = q[n-1-i], q[i]
	}
	return string(append(q, qs[dir]...))
}
