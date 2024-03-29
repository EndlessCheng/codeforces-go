// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	exampleIns := [][]string{{`[4,3,10,9,8]`}, {`[4,4,7,6,7]`}, {`[6]`}}
	exampleOuts := [][]string{{`[10,9]`}, {`[7,7,6]`}, {`[6]`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithCase(t, minSubsequence, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-183/problems/minimum-subsequence-in-non-increasing-order/
