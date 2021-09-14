package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	sampleIns := [][]string{{`[[1,0,1],[0,0,0],[1,0,1]]`}, {`[[1,0,0],[0,0,0],[0,0,0]]`}}
	sampleOuts := [][]string{{`2`}, {`4`}}
	if err := testutil.RunLeetCodeFunc(t, maxDistance, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
