// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,3,4,7,1,2,6]`, 
			`[1,3,4,1,2,6]`,
		},
		{
			`[1,2,3,4]`, 
			`[1,2,4]`,
		},
		{
			`[2,1]`, 
			`[2]`,
		},
		{
			`[1]`,
			`[]`,
		},
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, deleteMiddle, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-270/problems/delete-the-middle-node-of-a-linked-list/
