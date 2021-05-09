package main

// github.com/EndlessCheng/codeforces-go
func maximumPopulation(logs [][]int) (ans int) {
	c := [2050]int{}
	for _, p := range logs {
		for i := p[0]; i < p[1]; i++ {
			c[i]++
		}
	}
	for i := 1950; i < 2050; i++ {
		if c[i] > c[ans] {
			ans = i
		}
	}
	return
}
