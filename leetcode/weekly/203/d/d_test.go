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
			`[6,2,3,4,5,5]`, 
			`18`,
		},
		{
			`[7,7,7,7,7,7,7]`, 
			`28`,
		},
		{
			`[4]`, 
			`0`,
		},
		// TODO 测试参数的下界和上界
		{
			`[5,5]`,
			`5`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, stoneGameV, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-203/problems/stone-game-v/
