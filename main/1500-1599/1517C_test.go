package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1517/C
// https://codeforces.com/problemset/status/1517/problem/C
func TestCF1517C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 1
outputCopy
2
2 3
3 3 1
inputCopy
5
1 2 3 4 5
outputCopy
1
2 2
3 3 3
4 4 4 4
5 5 5 5 5`
	testutil.AssertEqualCase(t, rawText, 0, CF1517C)
}
