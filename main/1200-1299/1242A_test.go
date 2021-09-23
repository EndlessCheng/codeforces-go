package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1242/A
// https://codeforces.com/problemset/status/1242/problem/A
func TestCF1242A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
outputCopy
2
inputCopy
5
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1242A)
}
