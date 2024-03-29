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
			`3`, `[[1,3],[2,3]]`, `[3,2,5]`, 
			`8`,
		},
		{
			`5`, `[[1,5],[2,5],[3,5],[3,4],[4,5]]`, `[1,2,3,4,5]`, 
			`12`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimumTime, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-264/problems/parallel-courses-iii/
