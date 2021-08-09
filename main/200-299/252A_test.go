package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/252/A
// https://codeforces.com/problemset/status/252/problem/A
func TestCF252A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 1 1 2
outputCopy
3
inputCopy
3
1 2 7
outputCopy
7
inputCopy
4
4 2 4 8
outputCopy
14`
	testutil.AssertEqualCase(t, rawText, 0, CF252A)
}
