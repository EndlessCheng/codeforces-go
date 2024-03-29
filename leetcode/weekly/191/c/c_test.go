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
			`6`, `[[0,1],[1,3],[2,3],[4,0],[4,5]]`, 
			`3`, 
		},
		{
			`5`, `[[1,0],[1,2],[3,2],[3,4]]`, 
			`2`, 
		},
		{
			`3`, `[[1,0],[2,0]]`, 
			`0`, 
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minReorder, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-191/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
