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
			`[2,5]`, `[3,4]`, `2`, 
			`8`,
		},
		{
			`[-4,-2,0,3]`, `[2,4]`, `6`, 
			`0`,
		},
		{
			`[-2,-1,0,1,2]`, `[-3,-1,2,4,5]`, `3`, 
			`-6`,
		},
		
	}
	targetCaseNum := 1
	if err := testutil.RunLeetCodeFuncWithExamples(t, kthSmallestProduct, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-63/problems/kth-smallest-product-of-two-sorted-arrays/
