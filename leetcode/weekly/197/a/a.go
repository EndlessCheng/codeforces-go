package main

func numIdenticalPairs(a []int) (ans int) {
	for r, v := range a {
		for _, w := range a[:r] {
			if v == w {
				ans++
			}
		}
	}
	return
}
