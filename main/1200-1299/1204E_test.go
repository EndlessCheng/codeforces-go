package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1204/E
// https://codeforces.com/problemset/status/1204/problem/E
func TestCF1204E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0 2
outputCopy
0
inputCopy
2 0
outputCopy
2
inputCopy
2 2
outputCopy
5
inputCopy
2000 2000
outputCopy
674532367`
	testutil.AssertEqualCase(t, rawText, 0, CF1204E)
}
