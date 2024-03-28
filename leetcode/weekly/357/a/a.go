package main

import "slices"

// https://space.bilibili.com/206214
func finalString(s string) string {
	q := [2][]rune{}
	dir := 1
	for _, c := range s {
		if c == 'i' {
			dir ^= 1 // 修改添加方向
		} else {
			q[dir] = append(q[dir], c)
		}
	}
	slices.Reverse(q[dir^1])
	return string(append(q[dir^1], q[dir]...))
}
