package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1148/E
// https://codeforces.com/problemset/status/1148/problem/E
func TestCF1148E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 2 7 4 9
5 4 5 5 5
outputCopy
YES
4
4 3 1
2 3 1
2 5 2
1 5 2
inputCopy
3
1 5 10
3 5 7
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1148E)
}
