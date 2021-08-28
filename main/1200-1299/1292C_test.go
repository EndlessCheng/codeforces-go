package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1292/C
// https://codeforces.com/problemset/status/1292/problem/C
func TestCF1292C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2
2 3
outputCopy
3
inputCopy
5
1 2
1 3
1 4
3 5
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1292C)
}
