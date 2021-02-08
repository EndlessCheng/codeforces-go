package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1480/A
// https://codeforces.com/problemset/status/1480/problem/A
func TestCF1480A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
a
bbbb
az
outputCopy
b
azaz
by`
	testutil.AssertEqualCase(t, rawText, 0, CF1480A)
}
