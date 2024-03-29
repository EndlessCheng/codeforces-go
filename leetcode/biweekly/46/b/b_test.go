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
			`[[1,-1,-1],[3,-2,0]]`, `[1,-1,0,1,-1,-1,3,-2,0]`, 
			`true`,
		},
		{
			`[[10,-2],[1,2,3,4]]`, `[1,2,3,4,10,-2]`, 
			`false`,
		},
		{
			`[[1,2,3],[3,4]]`, `[7,7,1,2,3,4,7,7]`, 
			`false`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, canChoose, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-46/problems/form-array-by-concatenating-subarrays-of-another-array/
