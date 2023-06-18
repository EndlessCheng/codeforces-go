package main

// https://space.bilibili.com/206214
func distanceTraveled(mainTank, additionalTank int) (ans int) {
	for mainTank >= 5 {
		t := mainTank / 5
		ans += t * 50
		mainTank %= 5
		t = min(t, additionalTank)
		additionalTank -= t
		mainTank += t
	}
	return ans + mainTank*10
}

func min(a, b int) int { if b < a { return b }; return a }
