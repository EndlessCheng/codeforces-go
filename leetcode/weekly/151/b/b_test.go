package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	sampleIns := [][]string{{`["cbd"]`, `["zaaaz"]`}, {`["bbb","cc"]`, `["a","aa","aaa","aaaa"]`}}
	sampleOuts := [][]string{{`[1]`}, {`[1,2]`}}
	if err := testutil.RunLeetCodeFunc(t, numSmallerByFrequency, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
