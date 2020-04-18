package main

func minCount(a []int) (ans int) {
	for _, v := range a {
		ans += (v-1)/2 + 1
	}
	return
}
