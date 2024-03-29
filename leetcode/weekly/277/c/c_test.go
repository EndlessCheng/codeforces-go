// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	examples := [][]string{
		{
			`[10,6,5,8]`, 
			`[10,8]`,
		},
		{
			`[1,3,5,3]`, 
			`[1,5]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, findLonely, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-277/problems/find-all-lonely-numbers-in-the-array/
