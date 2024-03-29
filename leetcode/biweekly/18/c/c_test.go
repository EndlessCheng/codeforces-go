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
			`[[3,3,1,1],[2,2,1,2],[1,1,1,2]]`, 
			`[[1,1,1,1],[1,2,2,2],[1,2,3,3]]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, diagonalSort, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-18/problems/sort-the-matrix-diagonally/
