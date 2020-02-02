package main

import "sort"

func minSetSize(arr []int) (ans int) {
	cnts := make([]int, 1e5+1)
	for _, v := range arr {
		cnts[v]++
	}
	sort.Ints(cnts)
	sum := 0
	for i := len(cnts) - 1; i >= 0; i-- {
		ans++
		sum += cnts[i]
		if sum >= len(arr)/2 {
			break
		}
	}
	return
}
