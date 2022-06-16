package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/847/E
// https://codeforces.com/problemset/status/847/problem/E
func TestCF847E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
*..P*P*
outputCopy
3
inputCopy
10
.**PP.*P.*
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF847E)
}
