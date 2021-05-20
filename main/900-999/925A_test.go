package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/925/A
// https://codeforces.com/problemset/status/925/problem/A
func TestCF925A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6 1 1 3
2
5
3
1 1 5 6
1 3 5 4
3 3 5 3
outputCopy
7
5
4`
	testutil.AssertEqualCase(t, rawText, 0, CF925A)
}
