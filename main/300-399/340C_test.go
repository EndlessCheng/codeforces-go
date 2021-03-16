package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/340/C
// https://codeforces.com/problemset/status/340/problem/C
func TestCF340C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 5
outputCopy
22 3`
	testutil.AssertEqualCase(t, rawText, 0, CF340C)
}
