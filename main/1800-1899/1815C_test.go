package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1815/C
// https://codeforces.com/problemset/status/1815/problem/C
func TestCF1815C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 2
3 1
2 1
1 0
2 0
2 2
1 2
2 1
5 5
2 1
3 1
4 2
4 5
5 1
outputCopy
FINITE
5
2 3 1 2 3 
FINITE
1
1 
INFINITE
FINITE
3
2 1 2 
FINITE
10
4 2 3 5 4 1 3 2 5 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF1815C)
}
