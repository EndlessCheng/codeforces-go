package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/197/A
// https://codeforces.com/problemset/status/197/problem/A
func TestCF197A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5 2
outputCopy
First
inputCopy
6 7 4
outputCopy
Second`
	testutil.AssertEqualCase(t, rawText, 0, CF197A)
}
