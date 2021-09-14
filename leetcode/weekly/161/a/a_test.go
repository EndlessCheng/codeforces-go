package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`"xx"`, `"yy"`}, {`"xy"`, `"yx"`}, {`"xx"`, `"xy"`}, {`"xxyyxyxyxx"`, `"xyyxyxxxyx"`}}
	sampleOuts := [][]string{{`1`}, {`2`}, {`-1`}, {`4`}}
	if err := testutil.RunLeetCodeFunc(t, minimumSwap, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
