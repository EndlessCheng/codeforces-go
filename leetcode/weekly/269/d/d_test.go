// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`6`, `[[1,2,5],[2,3,8],[1,5,10]]`, `1`, 
			`[0,1,2,3,5]`,
		},
		{
			`4`, `[[3,1,3],[1,2,2],[0,3,3]]`, `3`, 
			`[0,1,3]`,
		},
		{
			`5`, `[[3,4,2],[1,2,1],[2,3,1]]`, `1`, 
			`[0,1,2,3,4]`,
		},
		{
			`6`, `[[0,2,1],[1,3,1],[4,5,1]]`, `1`, 
			`[0,1,2,3]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, findAllPeople, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-269/problems/find-all-people-with-secret/
