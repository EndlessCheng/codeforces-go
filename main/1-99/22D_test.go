package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/22/D
// https://codeforces.com/problemset/status/22/problem/D
func TestCF22D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
0 2
2 5
outputCopy
1
2 
inputCopy
5
0 3
4 2
4 8
8 10
7 7
outputCopy
3
7 10 3`
	testutil.AssertEqualCase(t, rawText, 0, CF22D)
}
