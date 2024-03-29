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
			`[1,4,8,10,20]`, `3`, 
			`5`,
		},
		{
			`[2,3,5,12,18]`, `2`, 
			`9`,
		},
		{
			`[7,4,6,1]`, `1`, 
			`8`,
		},
		{
			`[3,6,14,10]`, `4`, 
			`0`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minDistance, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-28/problems/allocate-mailboxes/
