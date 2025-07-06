package main

// https://space.bilibili.com/206214
func minMoves(sx, sy, x, y int) (ans int) {
	for ; x != sx || y != sy; ans++ {
		if x < sx || y < sy {
			return -1
		}
		if x == y {
			if sy > 0 {
				x = 0
			} else {
				y = 0
			}
			continue
		}
		// 保证 x > y
		if x < y {
			x, y = y, x
			sx, sy = sy, sx
		}
		if x >= y*2 {
			if x%2 > 0 {
				return -1
			}
			x /= 2
		} else {
			x -= y
		}
	}
	return
}
