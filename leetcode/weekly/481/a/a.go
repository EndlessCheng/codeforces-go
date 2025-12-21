package main

// https://space.bilibili.com/206214
func mirrorDistance(n int) int {
	rev := 0
	for x := n; x > 0; x /= 10 {
		rev = rev*10 + x%10
	}
	return abs(n - rev)
}

func abs(x int) int { if x < 0 { return -x }; return x }
