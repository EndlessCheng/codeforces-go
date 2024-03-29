// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[0,6,5,2,2,5,1,9,4]`, `1`, `2`, 
			`20`,
		},
		{
			`[3,8,1,3,2,1,8,9,0]`, `3`, `2`, 
			`29`,
		},
		{
			`[2,1,5,6,0,9,5,0,3,8]`, `4`, `3`, 
			`31`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxSumTwoNoOverlap, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-133/problems/maximum-sum-of-two-non-overlapping-subarrays/
