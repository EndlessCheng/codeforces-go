package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1517/D
// https://codeforces.com/problemset/status/1517/problem/D
func TestCF1517D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 10
1 1
1 1
1 1
1 1 1
1 1 1
outputCopy
10 10 10
10 10 10
10 10 10
inputCopy
2 2 4
1
3
4 2
outputCopy
4 4
10 6
inputCopy
2 2 3
1
2
3 4
outputCopy
-1 -1
-1 -1`
	testutil.AssertEqualCase(t, rawText, 0, CF1517D)
}
