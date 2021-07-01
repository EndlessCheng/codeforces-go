package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/478/D
// https://codeforces.com/problemset/status/478/problem/D
func TestCF478D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 6
outputCopy
2
inputCopy
9 7
outputCopy
6
inputCopy
1 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF478D)
}
