package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1468/A
// https://codeforces.com/problemset/status/1468/problem/A
func TestCF1468A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
8
1 2 7 3 2 1 2 3
2
2 1
7
4 1 5 2 6 3 7
outputCopy
6
2
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1468A)
}
