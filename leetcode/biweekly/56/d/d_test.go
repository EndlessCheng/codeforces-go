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
			`30`, `[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]`, `[5,1,2,20,20,3]`, 
			`11`,
		},
		{
			`29`, `[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]`, `[5,1,2,20,20,3]`, 
			`48`,
		},
		{
			`25`, `[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]`, `[5,1,2,20,20,3]`, 
			`-1`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minCost, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-56/problems/minimum-cost-to-reach-destination-in-time/
