// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	exampleIns := [][]string{{`["mass","as","hero","superhero"]`}, {`["leetcode","et","code"]`}, {`["blue","green","bu"]`}}
	exampleOuts := [][]string{{`["as","hero"]`}, {`["et","code"]`}, {`[]`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithCase(t, stringMatching, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
	if err := testutil.RunLeetCodeFuncWithCase(t, stringMatchingSA, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-184/problems/string-matching-in-an-array/
