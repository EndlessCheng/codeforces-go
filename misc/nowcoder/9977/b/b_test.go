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
			`"5"`,
			`false`,
		},
		// TODO 测试参数的下界和上界
		{
			`"1"`,
			`true`,
		},
		{
			`"0"`,
			`true`,
		},
		{
			`"2"`,
			`false`,
		},
		{
			`"31"`,
			`true`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, judge, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/9977/b
