package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	sampleIns := [][]string{{`["dog","cat","dad","good"]`, `["a","a","c","d","d","d","g","o","o"]`, `[1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]`}, {`["xxxz","ax","bx","cx"]`, `["z","a","b","c","x","x","x"]`, `[4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10]`}, {`["leetcode"]`, `["l","e","t","c","o","d"]`, `[0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0]`}}
	sampleOuts := [][]string{{`23`}, {`27`}, {`0`}}
	if err := testutil.RunLeetCodeFunc(t, maxScoreWords, sampleIns, sampleOuts); err != nil {
		t.Fatal(err)
	}
}
