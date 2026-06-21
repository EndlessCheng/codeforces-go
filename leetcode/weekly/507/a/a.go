package main

// https://space.bilibili.com/206214
func maxDistance(moves string) int {
	x, y, free := 0, 0, 0
	for _, ch := range moves {
		switch ch {
		case 'L': x--
		case 'R': x++
		case 'D': y--
		case 'U': y++
		default: free++
		}
	}
	return abs(x) + abs(y) + free
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
