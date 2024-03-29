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
			`[2,3,1,6,7]`, 
			`4`, 
		},
		{
			`[1,1,1,1,1]`, 
			`10`, 
		},
		{
			`[2,3]`, 
			`0`, 
		},
		{
			`[1,3,5,7,9]`, 
			`3`, 
		},
		{
			`[7,11,12,9,5,2,7,17,22]`, 
			`8`, 
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, countTriplets, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-188/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
