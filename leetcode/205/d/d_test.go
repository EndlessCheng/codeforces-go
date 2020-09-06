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
			`4`, `[[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]`, 
			`2`,
		},
		{
			`4`, `[[3,1,2],[3,2,3],[1,1,4],[2,1,4]]`, 
			`0`,
		},
		{
			`4`, `[[3,2,3],[1,1,2],[2,3,4]]`, 
			`-1`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 2
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxNumEdgesToRemove, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-205/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
