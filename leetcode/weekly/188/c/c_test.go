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
			`7`, `[[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]`, `[false,false,true,false,true,true,false]`, 
			`8`, 
		},
		{
			`7`, `[[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]`, `[false,false,true,false,false,true,false]`, 
			`6`, 
		},
		{
			`7`, `[[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]]`, `[false,false,false,false,false,false,false]`, 
			`0`, 
		},
		// TODO 测试参数的下界和上界
		{
			`4`,`[[0,1],[1,2],[0,3]]`,`[true,true,true,true]`,
			`6`,
		},
	}
	targetCaseNum := 4
	if err := testutil.RunLeetCodeFuncWithExamples(t, minTime, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-188/problems/minimum-time-to-collect-all-apples-in-a-tree/
