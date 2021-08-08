package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/805/B
// https://codeforces.com/problemset/status/805/problem/B
func TestCF805B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
aa
inputCopy
3
outputCopy
bba`
	testutil.AssertEqualCase(t, rawText, 0, CF805B)
}
