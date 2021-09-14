package main

// github.com/EndlessCheng/codeforces-go
func findGCD(a []int) int {
	min, max := a[0], a[0]
	for _, v := range a[1:] {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return gcd(min, max)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
