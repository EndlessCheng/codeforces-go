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
			`5`,`5`,`[[5,3],[2,1],[3,2],[3,3],[1,1],[3,1],[5,1],[4,2],[4,5],[1,3],[4,1],[3,5],[2,4],[2,5],[2,2],[1,4],[4,3],[5,2],[5,4],[5,5],[1,2],[1,5],[3,4],[2,3],[4,4]]`,
			`3`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, wwork, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/9476/c
