// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[3,2,4,3]`, `4`, `2`, 
			`[6,6]`,
		},
		{
			`[1,5,6]`, `3`, `4`, 
			`[2,3,2,2]`,
		},
		{
			`[1,2,3,4]`, `6`, `4`, 
			`[]`,
		},
		{
			`[1]`, `3`, `1`, 
			`[5]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, missingRolls, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-261/problems/find-missing-observations/
