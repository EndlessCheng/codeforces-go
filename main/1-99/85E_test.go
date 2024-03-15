package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/85/E
// https://codeforces.com/problemset/status/85/problem/E
func TestCF85E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
0 0
1 1
outputCopy
0
2
inputCopy
4
0 0
0 1
1 0
1 1
outputCopy
1
4
inputCopy
3
0 0
1000 1000
5000 5000
outputCopy
2000
2`
	testutil.AssertEqualCase(t, rawText, 0, CF85E)
}
