package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	sampleIns := [][]string{{`["cat","bt","hat","tree"]`, `"atach"`}, {`["hello","world","leetcode"]`, `"welldonehoneyr"`}}
	sampleOuts := [][]string{{`6`}, {`10`}}
	if err := testutil.RunLeetCodeFunc(t, countCharacters, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
