package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	sampleIns := [][]string{{`[[1,0],[0,1]]`}, {`[[1,0],[1,1]]`}, {`[[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]`}}
	sampleOuts := [][]string{{`0`}, {`3`}, {`4`}}
	if err := testutil.RunLeetCodeFunc(t, countServers, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
