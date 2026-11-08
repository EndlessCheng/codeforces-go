package main

// https://space.bilibili.com/206214
func countValidSubstrings(s string) (ans int64) {
	prev := make([][3]int, len(s)+1)
	last := [3]int{}
	for i, ch := range s {
		i++ // 下标从 1 开始，这样可以用 0 表示不存在
		last[ch-'a'] = i // 维护最近的每种字母的下标
		prev[i] = last
		left := prev[prev[last[2]][1]][0] // 最近的 'c' 的最近的 'b' 的最近的 'a' 的下标
		ans += int64(i - left) // 子串左端点最小值为 left+1
	}
	return
}
