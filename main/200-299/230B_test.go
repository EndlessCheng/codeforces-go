package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/230/B
// https://codeforces.com/problemset/status/230/problem/B
func TestCF230B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 5 6
outputCopy
YES
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF230B)
}
