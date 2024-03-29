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
			`1`, `3`, 
			`"c"`,
		},
		{
			`1`, `4`, 
			`""`,
		},
		{
			`3`, `9`, 
			`"cab"`,
		},
		{
			`2`, `7`, 
			`""`,
		},
		{
			`10`, `100`, 
			`"abacbabacb"`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, getHappyString, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-24/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
