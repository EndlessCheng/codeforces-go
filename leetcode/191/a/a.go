package main

func maxProduct(a []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, v := range a {
		for j, w := range a {
			if j == i {
				continue
			}
			ans = max(ans, (v-1)*(w-1))
		}
	}
	return
}
