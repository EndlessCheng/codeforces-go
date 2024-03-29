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
			`[1,2,3]`, 
			`4`,
		},
		{
			`[1,3,3]`, 
			`4`,
		},
		{
			`[4,-2,-3,4,1]`, 
			`59`,
		},
		
	}
	targetCaseNum :=3
	if err := testutil.RunLeetCodeFuncWithExamples(t, subArrayRanges, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-271/problems/sum-of-subarray-ranges/
