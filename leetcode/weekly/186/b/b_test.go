// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`[1,2,3,4,5,6,1]`, `3`}, {`[2,2,2]`, `2`}, {`[9,7,7,9,7,7,9]`, `7`}, {`[1,1000,1]`, `1`}, {`[1,79,80,1,1,1,200,1]`, `3`}}
	exampleOuts := [][]string{{`12`}, {`4`}, {`55`}, {`1`}, {`202`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithCase(t, maxScore, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-186/problems/maximum-points-you-can-obtain-from-cards/
