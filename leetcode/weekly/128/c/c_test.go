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
			`[1,2,3,4,5,6,7,8,9,10]`, `5`, 
			`15`,
		},
		{
			`[3,2,2,4,1,4]`, `3`, 
			`6`,
		},
		{
			`[1,2,3,1,1]`, `4`, 
			`3`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, shipWithinDays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-128/problems/capacity-to-ship-packages-within-d-days/
