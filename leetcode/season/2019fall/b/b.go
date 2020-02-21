package main

func fraction(exp []int) (ans []int) {
	n := len(exp)
	h := make([]int, n+2)
	h[0], h[1] = 0, 1
	k := make([]int, n+2)
	k[0], k[1] = 1, 0
	for i, a := range exp {
		h[i+2] = a*h[i+1] + h[i]
		k[i+2] = a*k[i+1] + k[i]
	}
	return []int{h[n+1], k[n+1]}
}
