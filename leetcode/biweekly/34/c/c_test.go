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
			`[1,2,3,10,4,2,3,5]`, 
			`3`,
		},
		{
			`[5,4,3,2,1]`, 
			`4`,
		},
		{
			`[1,2,3]`, 
			`0`,
		},
		{
			`[1]`, 
			`0`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, findLengthOfShortestSubarray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-34/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
