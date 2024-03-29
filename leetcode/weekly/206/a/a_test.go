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
			`[[1,0,0],
            [0,0,1],
            [1,0,0]]`,
			`1`,
		},
		{
			`[[1,0,0],
            [0,1,0],
            [0,0,1]]`,
			`3`,
		},
		{
			`[[0,0,0,1],
            [1,0,0,0],
            [0,1,1,0],
            [0,0,0,0]]`,
			`2`,
		},
		{
			`[[0,0,0,0,0],
            [1,0,0,0,0],
            [0,1,0,0,0],
            [0,0,1,0,0],
            [0,0,0,1,1]]`,
			`3`,
		},
		// TODO 测试参数的下界和上界

	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, numSpecial, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-206/problems/special-positions-in-a-binary-matrix/
