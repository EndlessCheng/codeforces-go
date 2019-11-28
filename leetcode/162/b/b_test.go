package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	sampleIns := [][]string{{`2`, `1`, `[1,1,1]`}, {`2`, `3`, `[2,2,1,1]`}, {`5`, `5`, `[2,1,2,0,1,0,1,2,0,1]`}}
	sampleOuts := [][]string{{`[[1,1,0],[0,0,1]]`}, {`[]`}, {`[[1,1,1,0,1,0,0,1,0,0],[1,0,1,0,0,0,1,1,0,1]]`}}
	if err := testutil.RunLeetCodeFunc(t, reconstructMatrix, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
