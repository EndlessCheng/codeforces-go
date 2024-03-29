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
			`3`, 
			`2`,
		},
		{
			`6`, 
			`9`,
		},
		// TODO 测试参数的下界和上界
		{
			`1`,
			`0`,
		},
		{
			`2`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minOperations, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-202/problems/minimum-operations-to-make-array-equal/
