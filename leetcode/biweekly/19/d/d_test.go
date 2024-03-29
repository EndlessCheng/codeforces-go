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
			`[100,-23,-23,404,100,23,23,23,3,404]`, 
			`3`,
		},
		{
			`[7]`, 
			`0`,
		},
		{
			`[7,6,9,6,9,6,9,7]`, 
			`1`,
		},
		{
			`[6,1,9]`, 
			`2`,
		},
		{
			`[11,22,7,7,7,7,7,7,7,22,13]`, 
			`3`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minJumps, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-19/problems/jump-game-iv/
