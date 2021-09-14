package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	sampleIns := [][]string{{`"abab"`}, {`"leetcode"`}}
	sampleOuts := [][]string{{`"bab"`}, {`"tcode"`}}
	if err := testutil.RunLeetCodeFunc(t, lastSubstring, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
