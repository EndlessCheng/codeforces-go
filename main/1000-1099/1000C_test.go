package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1000/C
// https://codeforces.com/problemset/status/1000/problem/C
func TestCF1000C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 3
1 3
3 8
outputCopy
6 2 1 
inputCopy
3
1 3
2 4
5 7
outputCopy
5 2 0 `
	testutil.AssertEqualCase(t, rawText, 0, CF1000C)
}
