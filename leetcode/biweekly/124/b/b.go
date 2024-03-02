package main

import "slices"

// https://space.bilibili.com/206214
func lastNonEmptyString(s string) string {
	var cnt, last [26]int
	for i, b := range s {
		b -= 'a'
		cnt[b]++
		last[b] = i
	}

	// 注：也可以再遍历一次字符串，但效率不如下面，毕竟至多 26 个数
	ids := []int{}
	mx := slices.Max(cnt[:])
	for i, c := range cnt {
		if c == mx {
			ids = append(ids, last[i])
		}
	}
	slices.Sort(ids)

	t := make([]byte, len(ids))
	for i, id := range ids {
		t[i] = s[id]
	}
	return string(t)
}
