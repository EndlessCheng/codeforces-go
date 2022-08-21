package main

// https://space.bilibili.com/206214
func shiftingLetters(s string, shifts [][]int) string {
	diff := make([]int, len(s)+1)
	for _, p := range shifts {
		x := p[2]*2 - 1 // 0 和 1 变成 -1 和 1
		diff[p[0]] += x
		diff[p[1]+1] -= x
	}
	t, shift := []byte(s), 0
	for i, c := range t {
		shift = (shift+diff[i])%26 + 26 // 防一手负数
		t[i] = (c-'a'+byte(shift))%26 + 'a'
	}
	return string(t)
}
