package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1442/D
// https://codeforces.com/problemset/status/1442/problem/D
func TestCF1442D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
2 5 10
3 1 2 3
2 1 20
outputCopy
26`
	testutil.AssertEqualCase(t, rawText, 0, CF1442D)
}
