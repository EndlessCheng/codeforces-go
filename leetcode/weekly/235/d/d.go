package main

// github.com/EndlessCheng/codeforces-go
func countDifferentSubsequenceGCDs(a []int) (ans int) {
	has := [2e5 + 1]bool{}
	for _, v := range a {
		has[v] = true
	}
	for i := 1; i <= 2e5; i++ {
		g := 0
		for j := i; j <= 2e5; j += i {
			if has[j] {
				g = gcd(g, j)
			}
		}
		if g == i {
			ans++
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
