package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/486/A
// https://codeforces.com/problemset/status/486/problem/A
func TestCF486A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
2
inputCopy
5
outputCopy
-3`
	testutil.AssertEqualCase(t, rawText, 0, CF486A)
}
