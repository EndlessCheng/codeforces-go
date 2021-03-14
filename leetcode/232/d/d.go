package main

// github.com/EndlessCheng/codeforces-go
func maximumScore(a []int, k int) (ans int) {
	const border int = -1
	type pair struct{ v, i int }

	n := len(a)
	posL := make([]int, n)
	stack := []pair{{border, -1}}
	for i, v := range a {
		for {
			if top := stack[len(stack)-1]; top.v < v {
				posL[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	posR := make([]int, n)
	stack = []pair{{border, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := stack[len(stack)-1]; top.v < v {
				posR[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	for i, v := range a {
		if l, r := posL[i]+1, posR[i]; l <= k && k < r {
			ans = max(ans, v*(r-l))
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
