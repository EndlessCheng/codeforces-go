package main

// github.com/EndlessCheng/codeforces-go
func closestCost(baseCosts []int, toppingCosts []int, target int) (ans int) {
	sum, minD := 0, int(1e9)
	var f func(int)
	f = func(p int) {
		if p == len(toppingCosts) {
			for _, cost := range baseCosts {
				s := sum + cost
				d := abs(target - s)
				if d < minD || d == minD && s < ans {
					ans = s
					minD = d
				}
			}
			return
		}
		f(p + 1)
		c := toppingCosts[p]
		for i := 0; i < 2; i++ {
			sum += c
			f(p + 1)
		}
		sum -= 2 * c
	}
	f(0)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
