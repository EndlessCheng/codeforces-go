package main

// f[k ~ n-1] = 1

// i in k-1..0
// f[i] = (f[i+1] + ... + f[i+maxPts]) / maxPts
// f[i-1] = (f[i] + ... + f[i+maxPts-1]) / maxPts

// https://space.bilibili.com/206214
func new21Game(n, k, maxPts int) float64 {
	n = min(n, k+maxPts-1)
	f := make([]float64, n+1)
	s := 0.0
	for i := n; i >= 0; i-- {
		if i >= k {
			f[i] = 1 // 初始值
		} else {
			f[i] = s / float64(maxPts)
		}
		// 为下个循环做准备
		// 下个循环计算的是 f[i] + ... + f[i+maxPts-1]，所以要把 f[i+maxPts] 减掉
		s += f[i]
		if i+maxPts <= n {
			s -= f[i+maxPts]
		}
	}
	return f[0]
}
