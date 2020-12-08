package main

// github.com/EndlessCheng/codeforces-go
type pair struct{ l, r int }

func minSumOfLengths(a []int, target int) int {
	const inf int = 1e9
	ans, minL, s, pos, ps := inf, inf, 0, map[int]int{0: 0}, []pair{}
	for i, v := range a {
		i++
		s += v
		if l, has := pos[s-target]; has {
			ps = append(ps, pair{l, i})
			for ps[0].r <= l {
				minL = min(minL, ps[0].r-ps[0].l)
				ps = ps[1:]
			}
			ans = min(ans, minL+i-l)
		}
		pos[s] = i
	}
	if ans < inf {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
