package main

// https://space.bilibili.com/206214
func minimumCost(s, t string, flipCost, swapCost, crossCost int) int64 {
	cnt := [2][2]int{}
	for i, ch := range s {
		cnt[ch&1][t[i]&1]++
	}

	a := cnt[0][1]
	b := cnt[1][0]
	if a > b {
		a, b = b, a
	}

	res1 := (a + b) * flipCost
	res2 := a*swapCost + (b-a)*flipCost
	avg := (a + b) / 2
	res3 := (avg-a)*crossCost + avg*swapCost + (a+b)%2*flipCost
	return int64(min(res1, res2, res3))
}
