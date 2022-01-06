package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1620/A
// https://codeforces.com/problemset/status/1620/problem/A
func TestCF1620A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
EEE
EN
ENNEENE
NENN
outputCopy
YES
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1620A)
}
