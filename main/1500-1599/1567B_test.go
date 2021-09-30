package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1567/B
// https://codeforces.com/problemset/status/1567/problem/B
func TestCF1567B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
2 1
2 0
1 10000
2 10000
outputCopy
3
2
3
2
3
inputCopy
5
208504 0
279723 121614
250307 294554
125503 125503
140893 140892
outputCopy
208504
279724
250308
125503
140893`
	testutil.AssertEqualCase(t, rawText, -1, CF1567B)
}
