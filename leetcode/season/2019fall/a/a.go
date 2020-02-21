package main

func game(a []int, b []int) (ans int) {
	for i, v := range a {
		if v == b[i] {
			ans++
		}
	}
	return
}
