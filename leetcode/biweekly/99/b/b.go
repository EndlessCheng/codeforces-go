package main

// https://space.bilibili.com/206214
func coloredCells(n int) int64 {
	return 1 + 2*int64(n)*int64(n-1)
}
