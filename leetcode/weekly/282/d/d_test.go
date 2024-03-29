// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	examples := [][]string{
		{
			`[[2,3],[3,4]]`, `5`, `4`, 
			`21`,
		},
		{
			`[[1,10],[2,2],[3,4]]`, `6`, `5`, 
			`25`,
		},
		
	}
	targetCaseNum := 1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimumFinishTime, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-282/problems/minimum-time-to-finish-the-race/
