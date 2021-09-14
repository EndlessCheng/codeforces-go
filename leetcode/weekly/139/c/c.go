package main

func addNegabinary(a1, a2 []int) []int {
	if len(a1) < len(a2) {
		a1, a2 = a2, a1
	}
	for i, j := len(a1)-1, len(a2)-1; j >= 0; {
		a1[i] += a2[j]
		i--
		j--
	}
	ans := append(make([]int, 2), a1...)
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] >= 2 {
			ans[i] -= 2
			if ans[i-1] >= 1 {
				ans[i-1]--
			} else {
				ans[i-1]++
				ans[i-2]++
			}
		}
	}
	for i, v := range ans {
		if v != 0 {
			return ans[i:]
		}
	}
	return []int{0}
}
