package main

// https://space.bilibili.com/206214
func minLights(lights []int) (ans int) {
	n := len(lights)
	diff := make([]int, n+1)
	for i, v := range lights {
		if v > 0 {
			// 照亮 [max(i-v, 0), min(i+v, n-1)]
			diff[max(i-v, 0)]++
			diff[min(i+v+1, n)]--
		}
	}

	sumD := 0
	for i, d := range diff[:n] {
		sumD += d
		if sumD == 0 {
			// 在 i+1 装一个灯泡，照亮 [i, i+2]
			ans++
			sumD++ // diff[i]++ 直接更新到 sumD 中
			diff[min(i+3, n)]--
		}
	}
	return
}
