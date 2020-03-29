package main

func findLucky(arr []int) (ans int) {
	cnts := [501]int{}
	for _, v := range arr {
		cnts[v]++
	}
	for i := 500; i > 0; i-- {
		if cnts[i] == i {
			return i
		}
	}
	return -1
}
