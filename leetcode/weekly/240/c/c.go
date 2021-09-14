package main

// github.com/EndlessCheng/codeforces-go
func maxSumMinProduct(a []int) (ans int) {
	type pair struct{ v, i int }
	n := len(a)
	sum := make([]int, n+1)
	posL := make([]int, n)
	stk := []pair{{0, -1}}
	for i, v := range a {
		sum[i+1] = sum[i] + v
		for {
			if top := stk[len(stk)-1]; top.v < v {
				posL[i] = top.i
				break
			}
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, pair{v, i})
	}
	posR := make([]int, n)
	stk = []pair{{0, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := stk[len(stk)-1]; top.v < v {
				posR[i] = top.i
				break
			}
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, pair{v, i})
	}

	for i, v := range a {
		ans = max(ans, v*(sum[posR[i]]-sum[posL[i]+1]))
	}
	return ans % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
