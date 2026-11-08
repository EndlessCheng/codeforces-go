package main

// https://space.bilibili.com/206214
func countValidSubstrings1(s string) (ans int64) {
	prev := make([][3]int, len(s)+1)
	last := [3]int{}
	for i, ch := range s {
		i++              // 下标从 1 开始，这样可以用 0 表示不存在
		last[ch-'a'] = i // 维护最近的每种字母的下标
		prev[i] = last
		left := prev[prev[last[2]][1]][0] // 最近的 'c' 的最近的 'b' 的最近的 'a' 的下标
		ans += int64(i - left)            // 子串左端点最小值为 left+1
	}
	return
}

func countValidSubstrings(s string) (ans int64) {
	lastA, lastAB, lastABC := -1, -1, -1
	for i, ch := range s {
		if ch == 'a' {
			lastA = i
		} else if ch == 'b' {
			lastAB = lastA
		} else {
			lastABC = lastAB
		}
		ans += int64(i - lastABC) // 子串左端点最小值为 lastABC+1
	}
	return
}
