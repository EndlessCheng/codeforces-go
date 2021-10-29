package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1004/C
// https://codeforces.com/problemset/status/1004/problem/C
func TestCF1004C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 5 4 1 3
outputCopy
9
inputCopy
7
1 2 1 1 1 3 2
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, -1, CF1004C)
}
