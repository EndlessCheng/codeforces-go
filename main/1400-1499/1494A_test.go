package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1494/A
// https://codeforces.com/problemset/status/1494/problem/A
func TestCF1494A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
AABBAC
CACA
BBBBAC
ABCA
outputCopy
YES
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1494A)
}
