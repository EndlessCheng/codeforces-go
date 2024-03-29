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
			`"B0B6G0R6R0R6G9"`, 
			`1`,
		},
		{
			`"B0R0G0R9R0B0G0"`, 
			`1`,
		},
		{
			`"G4"`, 
			`0`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, countPoints, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-271/problems/rings-and-rods/
