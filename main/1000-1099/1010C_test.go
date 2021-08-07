package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1010/C
// https://codeforces.com/problemset/status/1010/problem/C
func TestCF1010C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 8
12 20
outputCopy
2
0 4 
inputCopy
3 10
10 20 30
outputCopy
1
0 `
	testutil.AssertEqualCase(t, rawText, 0, CF1010C)
}
