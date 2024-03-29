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
			`[[1,2,3],[3,1,2],[2,3,1]]`, 
			`true`,
		},
		{
			`[[1,1,1],[1,2,3],[1,2,3]]`, 
			`false`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, checkValid, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-275/problems/check-if-every-row-and-column-contains-all-numbers/
