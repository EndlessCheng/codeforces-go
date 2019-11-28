package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	sampleIns := [][]string{{`[3,6,5,1,8]`}, {`[4]`}, {`[1,2,3,4,4]`}}
	sampleOuts := [][]string{{`18`}, {`0`}, {`12`}}
	if err := testutil.RunLeetCodeFunc(t, maxSumDivThree, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
