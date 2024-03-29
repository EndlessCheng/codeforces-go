// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	examples := [][]string{
		{
			`[5,3,6,1,12]`, `3`, 
			`24`,
		},
		{
			`[2,7,9]`, `4`, 
			`4`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, findFinalValue, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-278/problems/keep-multiplying-found-values-by-two/
