package main

func expectNumber(a []int) (ans int) {
	cnts := [1e6 + 1]int{}
	for _, v := range a {
		cnts[v]++
	}
	for _, c := range cnts {
		if c > 0 {
			ans++
		}
	}
	return
}
