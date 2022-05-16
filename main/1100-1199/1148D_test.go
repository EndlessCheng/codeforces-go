package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1148/D
// https://codeforces.com/problemset/status/1148/problem/D
func TestCF1148D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 7
6 4
2 10
9 8
3 5
outputCopy
3
1 5 3
inputCopy
3
5 4
3 2
6 1
outputCopy
3
3 2 1`
	testutil.AssertEqualCase(t, rawText, 0, CF1148D)
}
