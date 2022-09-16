package main

// https://space.bilibili.com/206214
func storedEnergy(storeLimit int, power []int, supply [][]int) (ans int) {
	j := 0
	for i, s := range supply {
		t := len(power)
		if i+1 < len(supply) {
			t = supply[i+1][0]
		}
		for min, max := s[1], s[2]; j < t; j++ {
			if p := power[j]; p < min {
				if ans -= min - p; ans < 0 {
					ans = 0
				}
			} else if p > max {
				if ans += p - max; ans > storeLimit {
					ans = storeLimit
				}
			}
		}
	}
	return
}
