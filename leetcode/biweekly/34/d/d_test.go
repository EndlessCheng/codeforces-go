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
			`[2,3,6,8,4]`, `1`, `3`, `5`, 
			`4`,
		},
		{
			`[4,3,1]`, `1`, `0`, `6`, 
			`5`,
		},
		{
			`[5,2,1]`, `0`, `2`, `3`, 
			`0`,
		},
		{
			`[2,1,5]`, `0`, `0`, `3`, 
			`2`,
		},
		{
			`[1,2,3]`, `0`, `2`, `40`, 
			`615088286`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, countRoutes, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-34/problems/count-all-possible-routes/
