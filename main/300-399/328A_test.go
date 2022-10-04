package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/328/problem/A
// https://codeforces.com/problemset/status/328/problem/A
func TestCF328A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
836 624 412 200
outputCopy
-12
inputCopy
1 334 667 1000
outputCopy
1333`
	testutil.AssertEqualCase(t, rawText, 0, CF328A)
}
