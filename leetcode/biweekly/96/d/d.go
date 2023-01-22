package main

// https://space.bilibili.com/206214
func isReachable(targetX, targetY int) bool {
	g := gcd(targetX, targetY)
	return g&(g-1) == 0
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
