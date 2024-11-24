package main

// https://space.bilibili.com/206214
func shiftDistance(s, t string, nextCost, previousCost []int) (ans int64) {
	const sigma = 26
	var nxtSum, preSum [sigma*2 + 1]int64
	for i := 0; i < sigma*2; i++ {
		nxtSum[i+1] = nxtSum[i] + int64(nextCost[i%sigma])
		preSum[i+1] = preSum[i] + int64(previousCost[i%sigma])
	}
	for i := range s {
		x := s[i] - 'a'
		y := t[i] - 'a'
		if y < x {
			y += sigma
		}
		res1 := nxtSum[y] - nxtSum[x]
		y = t[i] - 'a'
		if x < y {
			x += sigma
		}
		res2 := preSum[x+1] - preSum[y+1]
		ans += min(res1, res2)
	}
	return
}
