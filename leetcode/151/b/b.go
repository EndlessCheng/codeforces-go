package main

func numSmallerByFrequency(queries []string, words []string) (ans []int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	minW := make([]int, len(words))
	for i, w := range words {
		minC := int('z')
		for _, s := range w {
			minC = min(minC, int(s))
		}
		cnt := 0
		for _, s := range w {
			if int(s) == minC {
				cnt++
			}
		}
		minW[i] = cnt
	}

	ans = make([]int, len(queries))
	for i, q := range queries {
		minC := int('z')
		for _, s := range q {
			minC = min(minC, int(s))
		}
		cnt := 0
		for _, s := range q {
			if int(s) == minC {
				cnt++
			}
		}
		for _, cntw := range minW {
			if cnt < cntw {
				ans[i]++
			}
		}
	}
	return
}
