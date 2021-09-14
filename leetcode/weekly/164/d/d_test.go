package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	sampleIns := [][]string{{`3`, `2`}, {`2`, `4`}, {`4`, `2`}}
	sampleOuts := [][]string{{`4`}, {`2`}, {`8`}}
	if err := testutil.RunLeetCodeFunc(t, numWays, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
