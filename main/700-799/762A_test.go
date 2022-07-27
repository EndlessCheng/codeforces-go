package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/762/A
// https://codeforces.com/problemset/status/762/problem/A
func TestCF762A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
outputCopy
2
inputCopy
5 3
outputCopy
-1
inputCopy
12 5
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF762A)
}
