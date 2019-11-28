package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	sampleIns := [][]string{{`[1,1,2,1,1]`, `3`}, {`[2,4,6]`, `1`}, {`[2,2,2,1,2,2,1,2,2,2]`, `2`}}
	sampleOuts := [][]string{{`2`}, {`0`}, {`16`}}
	if err := testutil.RunLeetCodeFunc(t, numberOfSubarrays, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
