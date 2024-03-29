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
			`[[2,4],[6,8]]`, `2`, 
			`4`,
		},
		{
			`[[1,5],[2,3]]`, `1`, 
			`5`,
		},
		{
			`[[1,2],[3,4]]`, `2`, 
			`-1`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minOperations, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-262/problems/minimum-operations-to-make-a-uni-value-grid/
