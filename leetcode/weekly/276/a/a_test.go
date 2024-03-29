// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	examples := [][]string{
		{
			`"abcdefghi"`, `3`, `"x"`, 
			`["abc","def","ghi"]`,
		},
		{
			`"abcdefghij"`, `3`, `"x"`, 
			`["abc","def","ghi","jxx"]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, divideString, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-276/problems/divide-a-string-into-groups-of-size-k/
