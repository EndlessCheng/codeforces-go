package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/739/A
// https://codeforces.com/problemset/status/739/problem/A
func TestCF739A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1 3
2 5
4 5
outputCopy
2
1 0 2 1 0
inputCopy
4 2
1 4
2 4
outputCopy
3
5 2 0 1`
	testutil.AssertEqualCase(t, rawText, 0, CF739A)
}
