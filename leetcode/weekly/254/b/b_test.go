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
			`[1,2,3,4,5]`, 
			`[1,2,4,5,3]`,
		},
		{
			`[6,2,0,9,7]`, 
			`[9,7,6,2,0]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, rearrangeArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-254/problems/array-with-elements-not-equal-to-average-of-neighbors/
