package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func countCollisions(s string) int {
	s = strings.TrimLeft(s, "L")          // 前缀向左的车不会发生碰撞
	s = strings.TrimRight(s, "R")         // 后缀向右的车不会发生碰撞
	return len(s) - strings.Count(s, "S") // 剩下非停止的车必然会碰撞
}
