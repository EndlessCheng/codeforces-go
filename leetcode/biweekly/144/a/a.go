package main

import "math"

// https://space.bilibili.com/206214
func canAliceWin(n int) bool {
	x := (21 - int(math.Ceil(math.Sqrt(float64(441-n*8))))) / 2
	return x%2 > 0
}

func canAliceWin2(n int) bool {
	pick := 10
	for n >= pick {
		n -= pick
		pick--
	}
	return (10-pick)%2 > 0
}
