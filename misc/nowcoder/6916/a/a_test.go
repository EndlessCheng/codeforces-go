// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`5`,`2`,`[(1,2),(2,3),(3,4),(2,5)]`,
			`4`,
		},
		// TODO 测试参数的下界和上界

	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, solve, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/6916/a
