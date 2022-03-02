package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/909/B
// https://codeforces.com/problemset/status/909/problem/B
func TestCF909B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
2
inputCopy
3
outputCopy
4
inputCopy
4
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF909B)
}
