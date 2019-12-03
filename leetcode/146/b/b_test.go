package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	sampleIns := [][]string{{`3`, `[[0,1],[1,2]]`, `[]`}, {`3`, `[[0,1]]`, `[[2,1]]`}, {`3`, `[[1,0]]`, `[[2,1]]`}, {`3`, `[[0,1]]`, `[[1,2]]`}, {`3`, `[[0,1],[0,2]]`, `[[1,0]]`}}
	sampleOuts := [][]string{{`[0,1,-1]`}, {`[0,1,-1]`}, {`[0,-1,-1]`}, {`[0,1,2]`}, {`[0,1,1]`}}
	if err := testutil.RunLeetCodeFunc(t, shortestAlternatingPaths, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
