// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`3`,`[1,2]`,`[2,3]`,
			`3`,
		},
		// TODO 测试参数的下界和上界

	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, PointsOnDiameter, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/9753/c
