// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"slices"
	"testing"
)

func Test_d(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, maxRectangleArea, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-427/problems/maximum-area-rectangle-with-point-constraints-ii/
// https://leetcode.cn/problems/maximum-area-rectangle-with-point-constraints-ii/

func TestCompareInf(_t *testing.T) {
	return
	testutil.DebugTLE = 0
	rg := testutil2.NewRandGenerator()
	inputGenerator := func() (xs, ys []int) {
		//return
		rg.Clear()
		n := rg.Int(1, 5)
		xs = rg.IntSlice(n, 0, 5)
		ys = rg.IntSlice(n, 0, 5)
		return
	}

	testutil.CompareInf(_t, inputGenerator, maxRectangleAreaAC, maxRectangleArea)
}

func maxRectangleAreaAC(xCoord []int, yCoord []int) (ans int64) {
	ps := [][]int{}
	for i, x := range xCoord {
		ps = append(ps, []int{x, yCoord[i]})
	}
	slices.SortFunc(ps, func(a, b []int) int { return a[0] - b[0] })
	for i, p1 := range ps {
		for j, p2 := range ps {
			if j == i || p1[1] != p2[1] {
				continue
			}
			for k, p3 := range ps {
				if k == i || k == j || p3[0] != p1[0] {
					continue
				}
				low, up := p1[1], p3[1]
				if low > up {
					low, up = up, low
				}
			o:
				for l, p4 := range ps {
					if l == i || l == j || l == k || p4[0] != p2[0] || p4[1] != p3[1] {
						continue
					}
					for idx, p := range ps {
						if idx != i && idx != j && idx != k && idx != l {
							x, y := p[0], p[1]
							if p1[0] <= x && x <= p2[0] && low <= y && y <= up {
								continue o
							}
						}
					}
					ans = max(ans, int64((p2[0]-p1[0])*(up-low)))
				}
			}
		}
	}
	if ans == 0 {
		ans = -1
	}
	return
}
