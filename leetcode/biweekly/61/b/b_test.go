// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,3,4,2,6,8]`, 
			`[1,3,4]`,
		},
		{
			`[6,3,0,1]`, 
			`[]`,
		},
		{
			`[1]`, 
			`[]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, findOriginalArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-61/problems/find-original-array-from-doubled-array/
