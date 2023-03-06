package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/148/E
// https://codeforces.com/problemset/status/148/problem/E
func TestCF148E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
3 3 7 2
3 4 1 5
outputCopy
15
inputCopy
1 3
4 4 3 1 2
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF148E)
}
