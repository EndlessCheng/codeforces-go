// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	examples := [][]string{
		{
			`[11,7,2,15]`, 
			`2`,
		},
		{
			`[-3,3,3,90]`, 
			`2`,
		},
		{
			`[723,723,-423,723,-647,532,723,723,212,-391,723]`,
			`4`,
		},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, countElements, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-277/problems/count-elements-with-strictly-smaller-and-greater-elements/
