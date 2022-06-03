package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/949/A
// https://codeforces.com/problemset/status/949/problem/A
func TestCF949A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0010100
outputCopy
3
3 1 3 4
3 2 5 6
1 7
inputCopy
111
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF949A)
}
