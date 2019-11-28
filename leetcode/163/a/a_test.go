package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`[[1,2,3],[4,5,6],[7,8,9]]`, `1`}, {`[[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]]`, `4`}, {`[[1,2,3],[4,5,6],[7,8,9]]`, `9`}}
	sampleOuts := [][]string{{`[[9,1,2],[3,4,5],[6,7,8]]`}, {`[[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]`}, {`[[1,2,3],[4,5,6],[7,8,9]]`}}
	if err := testutil.RunLeetCodeFunc(t, shiftGrid, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
