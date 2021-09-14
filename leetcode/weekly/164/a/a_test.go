package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`[[1,1],[3,4],[-1,0]]`}, {`[[3,2],[-2,2]]`}}
	sampleOuts := [][]string{{`7`}, {`5`}}
	if err := testutil.RunLeetCodeFunc(t, minTimeToVisitAllPoints, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
