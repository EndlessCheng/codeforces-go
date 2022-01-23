package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1594/D
// https://codeforces.com/problemset/status/1594/problem/D
func TestCF1594D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 2
1 2 imposter
2 3 crewmate
5 4
1 3 crewmate
2 5 crewmate
2 4 imposter
3 4 imposter
2 2
1 2 imposter
2 1 crewmate
3 5
1 2 imposter
1 2 imposter
3 2 crewmate
3 2 crewmate
1 3 imposter
5 0
outputCopy
2
4
-1
2
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1594D)
}
