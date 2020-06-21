// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[1,2,3,4]`, 
			`[-1,-1,-1,-1]`,
		},
		{
			`[1,2,0,0,2,1]`, 
			`[-1,-1,2,1,-1,-1]`,
		},
		{
			`[1,2,0,1,2]`, 
			`[]`,
		},
		{
			`[69,0,0,0,69]`, 
			`[-1,69,1,1,-1]`,
		},
		{
			`[10,20,20]`, 
			`[]`,
		},
		// TODO 测试参数的下界和上界
		{
			`[0,1,1]`,
			`[]`,
		},
		{
			`[1,0,2,3,0,1,2]`,
			`[-1,1,-1,-1,2,-1,-1]`,
		},
		{
			`[1,0,2,3,0,2,1]`,
			`[-1,1,-1,-1,2,-1,-1]`,
		},
		{
			`[1,2,0,2,3,0,1]`,
			`[-1,-1,2,-1,-1,1,-1]`,
		},
		{
			`[3,5,4,0,1,0,1,5,2,8,9]`,
			`[-1,-1,-1,5,-1,1,-1,-1,-1,-1,-1]`,
		},
		{
			`[2,3,0,0,3,1,0,1,0,2,2]`,
			`[]`,
		},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, avoidFlood, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-194/problems/avoid-flood-in-the-city/
