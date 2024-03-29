// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`7`, `[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]`, `2`, `4`}, {`7`, `[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]`, `1`, `7`}, {`7`, `[[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]`, `20`, `6`}}
	exampleOuts := [][]string{{`0.16666666666666666`}, {`0.3333333333333333`}, {`0.16666666666666666`}}
	// TODO: 测试参数的下界和上界！
	// custom test cases or WA cases.
	exampleIns = append(exampleIns,
		[]string{`9`, `[[2,1],[3,1],[4,2],[5,3],[6,5],[7,4],[8,7],[9,7]]`, `1`, `8`},
	)
	exampleOuts = append(exampleOuts,
		[]string{`0`},
	)
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithCase(t, frogPosition, exampleIns, exampleOuts, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode-cn.com/contest/weekly-contest-179/problems/frog-position-after-t-seconds/

//9
//[[2,1],[3,1],[4,2],[5,3],[6,5],[7,4],[8,7],[9,7]]
//1
//8
