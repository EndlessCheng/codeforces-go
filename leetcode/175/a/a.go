package main

func checkIfExist(arr []int) (ans bool) {
	cnts := map[int]int{}
	for _, v := range arr {
		cnts[v]++
	}
	if cnts[0] >= 2 {
		return true
	}
	delete(cnts, 0)
	for k := range cnts {
		if cnts[2*k] > 0 {
			return true
		}
	}
	return
}
