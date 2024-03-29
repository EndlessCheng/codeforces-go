// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[0,0,0,0,0]`, `[[1,10],[10,1],[10,1],[1,10],[5,1]]`, `5`, `2`, `3`, 
			`9`,
		},
		{
			`[0,2,1,2,0]`, `[[1,10],[10,1],[10,1],[1,10],[5,1]]`, `5`, `2`, `3`, 
			`11`,
		},
		{
			`[0,0,0,0,0]`, `[[1,10],[10,1],[1,10],[10,1],[1,10]]`, `5`, `2`, `5`, 
			`5`,
		},
		{
			`[3,1,2,3]`, `[[1,1,1],[1,1,1],[1,1,1],[1,1,1]]`, `4`, `3`, `3`, 
			`-1`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minCost, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-192/problems/paint-house-iii/
