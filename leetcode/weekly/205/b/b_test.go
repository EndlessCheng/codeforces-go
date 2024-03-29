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
			`[7,4]`, `[5,2,8,9]`, 
			`1`,
		},
		{
			`[1,1]`, `[1,1,1]`, 
			`9`,
		},
		{
			`[7,7,8,3]`, `[1,2,9,7]`, 
			`2`,
		},
		{
			`[4,7,9,11,23]`, `[3,5,1024,12,18]`, 
			`0`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numTriplets, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-205/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/
