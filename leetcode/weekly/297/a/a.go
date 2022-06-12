package main

// https://space.bilibili.com/206214/dynamic
func calculateTax(brackets [][]int, income int) float64 {
	ans, pre := 0, 0
	for _, b := range brackets {
		up, p := b[0], b[1]
		if income <= up {
			ans += (income - pre) * p
			break
		}
		ans += (up - pre) * p
		pre = up
	}
	return float64(ans) / 100
}
