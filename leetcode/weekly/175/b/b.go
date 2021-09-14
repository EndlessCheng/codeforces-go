package main

func minSteps(s string, t string) (ans int) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	cnts := [26]int{}
	for _, b := range s {
		b -= 'a'
		cnts[b]++
	}
	for _, b := range t {
		b -= 'a'
		cnts[b]--
	}
	for _, c := range cnts {
		ans += abs(c)
	}
	return ans / 2
}
