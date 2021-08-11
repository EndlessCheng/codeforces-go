package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1225/E
// https://codeforces.com/problemset/status/1225/problem/E
func TestCF1225E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
.
outputCopy
1
inputCopy
2 3
...
..R
outputCopy
0
inputCopy
4 4
...R
.RR.
.RR.
R...
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1225E)
}
