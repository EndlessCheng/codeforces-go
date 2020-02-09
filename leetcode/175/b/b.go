package main

func minSteps(s string, t string) (ans int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	cntS := [26]int{}
	for _, b := range s {
		b -= 'a'
		cntS[b]++
	}
	cntT := [26]int{}
	for _, b := range t {
		b -= 'a'
		cntT[b]++
	}
	for i, c := range cntS {
		ans += abs(c - cntT[i])
	}
	return ans / 2
}
