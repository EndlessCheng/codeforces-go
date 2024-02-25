package main

// https://space.bilibili.com/206214
func largestSquareArea(bottomLeft, topRight [][]int) (ans int64) {
	for i, b1 := range bottomLeft {
		t1 := topRight[i]
		for j := i + 1; j < len(bottomLeft); j++ {
			b2, t2 := bottomLeft[j], topRight[j]
			height := min(t1[1], t2[1]) - max(b1[1], b2[1])
			width := min(t1[0], t2[0]) - max(b1[0], b2[0])
			size := min(width, height)
			if size > 0 {
				ans = max(ans, int64(size)*int64(size))
			}
		}
	}
	return
}
