// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, minCost, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-316/problems/minimum-cost-to-make-array-equal/

func TestCompareInf(t *testing.T) {
	testutil.DebugTLE = 0

	inputGenerator := func() (nums, cost []int) {
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 9)
		nums = rg.IntSlice(n, 1, 9)
		cost = rg.IntSlice(n, 1, 1)
		return
	}

	testutil.CompareInf(t, inputGenerator, minCost2, minCost)
}
