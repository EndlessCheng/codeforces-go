package main

// https://space.bilibili.com/206214
func mostCompetitive(a []int, k int) (st []int) {
	k = len(a) - k
	for _, v := range a {
		for len(st) > 0 && k > 0 && v < st[len(st)-1] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, v)
	}
	st = st[:len(st)-k]
	return
}
