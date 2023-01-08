package main

// https://space.bilibili.com/206214
func categorizeBox(length, width, height, mass int) string {
	x := length >= 1e4 || width >= 1e4 || height >= 1e4 || length*width*height >= 1e9
	y := mass >= 100
	switch {
	case x && y: return "Both"
	case x: return "Bulky"
	case y: return "Heavy"
	default: return "Neither"
	}
}
