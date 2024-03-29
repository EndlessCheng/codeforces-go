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
			`3`, `[-3,-2,-1,0,0,1,2,3]`, 
			`[1,2,-3]`,
		},
		{
			`2`, `[0,0,0,0]`, 
			`[0,0]`,
		},
		{
			`4`, `[0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8]`, 
			`[0,-1,4,5]`,
		},
		{
			`1`, `[0,1]`,
			`[1]`,
		},
		{
			`1`, `[0,-1]`,
			`[-1]`,
		},
		{
			`3`, `[-574,-394,0,180,-180,-574,-754,0]`,
			`[180,-180,-574]`,
		},
		{
			`4`, `[305,-76,-381,0,-457,-183,-762,-381,503,198,884,579,198,122,-76,503]`,
			`[305,-381,-381,579]`,
		},
	}
	targetCaseNum := 2
	if err := testutil.RunLeetCodeFuncWithExamples(t, recoverArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-255/problems/find-array-given-subset-sums/
