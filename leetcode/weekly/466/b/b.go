package main

// https://space.bilibili.com/206214
func minOperations(s string) int {
	minC := 'z' + 1
	for _, c := range s {
		if c != 'a' {
			minC = min(minC, c)
		}
	}
	return int('z' + 1 - minC)
}
