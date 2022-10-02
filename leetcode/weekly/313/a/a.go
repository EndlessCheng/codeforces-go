package main

// https://space.bilibili.com/206214
func commonFactors(a, b int) (ans int) {
	g := gcd(a, b)
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i < g {
				ans++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
