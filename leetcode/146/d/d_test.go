package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	sampleIns := [][]string{{`[1,2,3,4]`, `[-1,4,5,6]`}, {`[1,-2,-5,0,10]`, `[0,-2,-1,-7,-4]`}}
	sampleOuts := [][]string{{`13`}, {`20`}}
	if err := testutil.RunLeetCodeFunc(t, maxAbsValExpr, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
