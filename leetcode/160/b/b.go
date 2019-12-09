package main

func circularPermutation(n int, start int) (ans []int) {
	if n == 1 {
		if start == 0 {
			return []int{0, 1}
		}
		return []int{1, 0}
	}

	ans0 := circularPermutation(n-1, 0)
	ans1 := make([]int, len(ans0))
	for i, v := range ans0 {
		ans1[len(ans0)-i-1] = v
	}
	//last := ans0[len(ans0)-1]
	//ans1 := circularPermutation(n-1, last)
	for i, v := range ans1 {
		ans1[i] = v | 1<<uint(n-1)
	}
	ans = append(ans0, ans1...)
	var pos int
	for i, v := range ans {
		if v == start {
			pos = i
			break
		}
	}
	ans = append(ans[pos:], ans[:pos]...)
	return
}
