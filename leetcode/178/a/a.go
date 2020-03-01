package main

func smallerNumbersThanCurrent(a []int) (ans []int) {
	ans = make([]int, len(a))
	for i, v := range a {
		for _, w := range a {
			if w < v {
				ans[i]++
			}
		}
	}
	return
}
