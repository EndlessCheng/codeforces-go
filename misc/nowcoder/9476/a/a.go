package main

//import . "nc_tools"
import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func doll(n, _ int, ps []*Interval) int {
	sort.Slice(ps, func(i, j int) bool { return ps[i].Start < ps[j].Start })
	return sort.Search(3e9, func(d int) bool {
		cnt, pre := 0, int(-1e10)
		for _, p := range ps {
			l, r := p.Start, p.End
			if l < pre+d {
				l = pre + d
			}
			if l > r {
				continue
			}
			c := (r - l) / d
			cnt += 1 + c
			pre = l + c*d
		}
		return cnt < n
	}) - 1
}
