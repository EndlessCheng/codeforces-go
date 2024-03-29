// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`["S#O", "M..", "M.T"]`,
			`16`,
		},
		{
			`["S#O", "M.#", "M.T"]`,
			`-1`,
		},
		{
			`["S#O", "M.T", "M.."]`,
			`17`,
		},
		// TODO: 测试参数的下界和上界！
		{
			`["...O.",".S#M.","..#T.","....."]`,
			`5`,
		},
		{
			`
[
"#####.", 
"O##..T", 
"M#.#.#", 
"OM.###", 
"#.O###", 
".MO##O", 
"#M..MO", 
"###O..", 
"##M.O#", 
"###..S"]`,
			`-1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimalSteps, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/season/2020-spring/problems/xun-bao/
