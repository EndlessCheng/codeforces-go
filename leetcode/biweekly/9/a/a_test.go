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
			`[100,200,150,1000]`, 
			`4`,
		},
		{
			`[900,950,800,1000,700,800]`, 
			`5`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxNumberOfApples, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-9/problems/how-many-apples-can-you-put-into-the-basket/
