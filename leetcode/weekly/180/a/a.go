package main

func luckyNumbers(a [][]int) (ans []int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for _, row := range a {
		for j, v := range row {
			mi := int(1e9)
			for _, v := range row {
				mi = min(mi, v)
			}
			mx := 0
			for _, row := range a {
				mx = max(mx, row[j])
			}
			if v == mi && v == mx {
				ans = append(ans, v)
			}
		}
	}
	return
}
