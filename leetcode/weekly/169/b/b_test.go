// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`[2,1,4]`, `[1,0,3]`}, {`[0,-10,10]`, `[5,1,7,0,2]`}, {`[]`, `[5,1,7,0,2]`}, {`[0,-10,10]`, `[]`}, {`[1,null,8]`, `[8,1]`}}
	exampleOuts := [][]string{{`[0,1,1,2,3,4]`}, {`[-10,0,0,1,2,5,7,10]`}, {`[0,1,2,5,7]`}, {`[-10,0,10]`}, {`[1,1,8,8]`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	const targetCaseNum = 0
	if err := testutil.RunLeetCodeFuncWithCase(t, getAllElements, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
