package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/414/A
// https://codeforces.com/problemset/status/414/problem/A
func TestCF414A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
outputCopy
1 2 3 4 5
inputCopy
5 3
outputCopy
2 4 3 7 1
inputCopy
7 2
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF414A)
}
