package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	exampleIns := [][]string{{`[1,2,3,2,null,2,4]`, `2`}, {`[1,3,3,3,2]`, `3`}, {`[1,2,null,2,null,2]`, `2`}, {`[1,1,1]`, `1`}, {`[1,2,3]`, `1`}}
	exampleOuts := [][]string{{`[1,null,3,null,4]`}, {`[1,3,null,null,2]`}, {`[1]`}, {`[]`}, {`[1,2,3]`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	const targetCaseNum = 0
	if err := testutil.RunLeetCodeFuncWithCase(t, removeLeafNodes, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
