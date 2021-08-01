package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1234/F
// https://codeforces.com/problemset/status/1234/problem/F
func TestCF1234F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abacaba
outputCopy
3
inputCopy
abcdecdf
outputCopy
6
inputCopy
aabbcc
outputCopy
3
inputCopy
abcdeefc
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1234F)
}
