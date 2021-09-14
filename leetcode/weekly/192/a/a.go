package main

func shuffle(a []int, n int) (ans []int) {
	for i, v := range a[:n] {
		ans = append(ans, v, a[i+n])
	}
	return
}
