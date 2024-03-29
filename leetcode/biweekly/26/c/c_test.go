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
			`[3,1,4,3,null,1,5]`, 
			`4`,
		},
		{
			`[3,3,null,4,2]`, 
			`3`,
		},
		{
			`[1]`, 
			`1`,
		},
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, goodNodes, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-26/problems/count-good-nodes-in-binary-tree/
