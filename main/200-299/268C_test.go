package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/268/C
// https://codeforces.com/problemset/status/268/problem/C
func TestCF268C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
outputCopy
3
0 1
1 2
2 0
inputCopy
4 3
outputCopy
4
0 3
2 1
3 0
4 2`
	testutil.AssertEqualCase(t, rawText, 0, CF268C)
}
