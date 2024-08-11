package main

// https://space.bilibili.com/206214
func finalPositionOfSnake(n int, commands []string) int {
	i, j := 0, 0
	for _, s := range commands {
		switch s[0] {
		case 'U': i--
		case 'D': i++
		case 'L': j--
		default:  j++
		}
	}
	return i*n + j
}
