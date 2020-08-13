package main

import (
	//. "nc_tools"
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// github.com/EndlessCheng/codeforces-go
func solve(n, _ int, ps []*Point) int {
	invalid := [21][21]bool{}
	for _, p := range ps {
		invalid[p.X][p.Y] = true
		invalid[p.Y][p.X] = true
	}
	cnt := 0
	chosen := []int{}
	var f func(int)
	f = func(p int) {
		if p == n+1 {
			cnt++
			return
		}
		f(p + 1)
		for _, v := range chosen {
			if invalid[v][p] {
				return
			}
		}
		chosen = append(chosen, p)
		f(p + 1)
		chosen = chosen[:len(chosen)-1]
	}
	f(1)
	return cnt
}
