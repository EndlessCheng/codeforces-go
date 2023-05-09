package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1118/D2
// https://codeforces.com/problemset/status/1118/problem/D2
func TestCF1118D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 8
2 3 1 1 2
outputCopy
4
inputCopy
7 10
1 3 4 2 1 4 2
outputCopy
2
inputCopy
5 15
5 5 5 5 5
outputCopy
1
inputCopy
5 16
5 5 5 5 5
outputCopy
2
inputCopy
5 26
5 5 5 5 5
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1118D2)
}
