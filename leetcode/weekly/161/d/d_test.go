package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	sampleIns := [][]string{{`[12,5,7,23]`}, {`[29,6,10]`}, {`[3,6]`}}
	sampleOuts := [][]string{{`true`}, {`true`}, {`false`}}
	if err := testutil.RunLeetCodeFunc(t, isGoodArray, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
