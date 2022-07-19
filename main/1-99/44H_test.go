package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/44/H
// https://codeforces.com/problemset/status/44/problem/H
func TestCF44H(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
12345
outputCopy
48
inputCopy
09
outputCopy
15`
	testutil.AssertEqualCase(t, rawText, 0, CF44H)
}
