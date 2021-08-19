package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1182/E
// https://codeforces.com/problemset/status/1182/problem/E
func TestCF1182E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 1 2 5 3
outputCopy
72900
inputCopy
17 97 41 37 11
outputCopy
317451037`
	testutil.AssertEqualCase(t, rawText, 0, CF1182E)
}
