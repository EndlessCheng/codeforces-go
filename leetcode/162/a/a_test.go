package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`2`, `3`, `[[0,1],[1,1]]`}, {`2`, `2`, `[[1,1],[0,0]]`}}
	sampleOuts := [][]string{{`6`}, {`0`}}
	if err := testutil.RunLeetCodeFunc(t, oddCells, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
