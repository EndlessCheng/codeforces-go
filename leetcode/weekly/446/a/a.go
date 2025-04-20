package main

// https://space.bilibili.com/206214
func calculateScore(instructions []string, values []int) (ans int64) {
	n := len(instructions)
	i := 0
	for 0 <= i && i < n && instructions[i] != "" {
		s := instructions[i]
		instructions[i] = ""
		if s[0] == 'a' {
			ans += int64(values[i])
			i++
		} else {
			i += values[i]
		}
	}
	return
}
