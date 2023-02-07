package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1141/F2
// https://codeforces.com/problemset/status/1141/problem/F2
func TestCF1141F2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
4 1 2 2 1 5 3
outputCopy
3
7 7
2 3
4 5
inputCopy
11
-5 -4 -3 -2 -1 0 1 2 3 4 5
outputCopy
2
3 4
1 1
inputCopy
4
1 1 1 1
outputCopy
4
4 4
1 1
2 2
3 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1141F2)
}
