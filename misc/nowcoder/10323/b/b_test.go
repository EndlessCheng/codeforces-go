// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`1`,
			`3`,
		},
		{
			`2`,
			`6`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, wwork, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/10323/b
