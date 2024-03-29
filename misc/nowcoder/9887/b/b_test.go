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
			`"ABA"`,
			`1`,
		},
		{
			`"ABC"`,
			`-1`,
		},
		// TODO 测试参数的下界和上界
		{
			`"ABABA"`,
			`3`,
		},
		{
			`"AAA"`,
			`2`,
		},
		{
			`"A"`,
			`-1`,
		},
		{
			`"AAAA"`,
			`3`,
		},
		{
			`"ABAB"`,
			`2`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, solve, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/9887/b
