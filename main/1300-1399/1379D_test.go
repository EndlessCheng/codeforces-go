package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1379/D
// https://codeforces.com/problemset/status/1379/problem/D
func TestCF1379D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 24 60 15
16 0
17 15
outputCopy
0 0

inputCopy
2 24 60 16
16 0
17 15
outputCopy
1 0
2 `
	testutil.AssertEqualCase(t, rawText, 0, CF1379D)
}
