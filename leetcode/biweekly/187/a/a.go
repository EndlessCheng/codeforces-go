package main

import "slices"

// https://space.bilibili.com/206214
func rearrangeString1(s string, x, y byte) string {
	t := []byte(s)
	if x < y {
		slices.SortFunc(t, func(a, b byte) int { return int(b) - int(a) })
	} else {
		slices.Sort(t)
	}
	return string(t)
}

func rearrangeString(s string, x, y byte) string {
	t := []byte(s)
	l, r := 0, len(t)-1
	for l < r { // 循环直到不足两个字母
		if t[l] != x { // 寻找最左边的 x
			l++
		} else if t[r] != y { // 寻找最右边的 y
			r--
		} else {
			t[l], t[r] = t[r], t[l]
			// 交换后，[0,l] 都不含 x，[r,n-1] 都不含 y，问题变成 [l+1,r-1] 的子问题
			l++
			r--
		}
	}
	return string(t)
}
