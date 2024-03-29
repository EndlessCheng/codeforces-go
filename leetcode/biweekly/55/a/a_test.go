// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[1,2,10,5,7]`, 
			`true`,
		},
		{
			`[2,3,1,2]`, 
			`false`,
		},
		{
			`[1,1,1]`, 
			`false`,
		},
		{
			`[1,2,3]`, 
			`true`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, canBeIncreasing, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-55/problems/remove-one-element-to-make-the-array-strictly-increasing/
