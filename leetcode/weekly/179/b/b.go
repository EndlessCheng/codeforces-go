package main

func numTimesAllBlue(light []int) (ans int) {
	n := len(light)
	vis := make([]bool, n)
	j := 0
	for i, pos := range light {
		vis[pos-1] = true
		for ; j < n && vis[j]; j++ {
		}
		if j-1 == i {
			ans++
		}
	}
	return
}
