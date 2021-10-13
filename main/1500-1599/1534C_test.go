package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1534/C
// https://codeforces.com/problemset/status/1534/problem/C
func TestCF1534C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4
1 4 2 3
3 2 1 4
8
2 6 5 1 4 3 7 8
3 8 7 5 1 2 4 6
outputCopy
2
8`
	testutil.AssertEqualCase(t, rawText, 0, CF1534C)
}
