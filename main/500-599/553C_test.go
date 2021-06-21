package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/553/C
// https://codeforces.com/problemset/status/553/problem/C
func TestCF553C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 0
outputCopy
4
inputCopy
4 4
1 2 1
2 3 1
3 4 0
4 1 0
outputCopy
1
inputCopy
4 4
1 2 1
2 3 1
3 4 0
4 1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 2, CF553C)
}
