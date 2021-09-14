package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`[[1,2],[2,1],[3,4],[5,6]]`}}
	sampleOuts := [][]string{{`1`}}
	if err := testutil.RunLeetCodeFunc(t, numEquivDominoPairs, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
